package httpclient

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptrace"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// 输出连接日志 1 开启 0 关闭
var OpenConnlog int32

const (
	MaxIdleConns        int = 100
	MaxIdleConnsPerHost int = 100
	IdleConnTimeout     int = 90
)

var defaultSetting = HttpSettings{
	UserAgent:        "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.89 Safari/537.36",
	ConnectTimeout:   60 * time.Second,
	ReadWriteTimeout: 60 * time.Second,
	Gzip:             true,
	DumpBody:         true,
	TlsClientConfig:  &tls.Config{InsecureSkipVerify: true}, //default ignore cer check
	EnableCookie:     true,
	ManualSetCookie:  false,
	Transport: &http.Transport{
		MaxConnsPerHost:       20,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
	// Proxy: func(req *http.Request) (*url.URL, error) {
	// 	u, _ := url.ParseRequestURI("http://127.0.0.1:8888")
	// 	return u, nil
	// },
}

var defaultCookieJar http.CookieJar
var settingMutex sync.Mutex

// createDefaultCookie creates a global cookiejar to store cookies.
func createDefaultCookie() {
	settingMutex.Lock()
	defer settingMutex.Unlock()
	defaultCookieJar, _ = cookiejar.New(nil)
}

// Overwrite default settings
func SetDefaultSetting(setting HttpSettings) {
	settingMutex.Lock()
	defer settingMutex.Unlock()
	defaultSetting = setting
	if defaultSetting.ConnectTimeout == 0 {
		defaultSetting.ConnectTimeout = 60 * time.Second
	}
	if defaultSetting.ReadWriteTimeout == 0 {
		defaultSetting.ReadWriteTimeout = 60 * time.Second
	}
	if defaultSetting.EnableCookie {
		defaultSetting.Cookies, _ = cookiejar.New(nil)
	}
}

// return *HttpRequest with specific method
func NewRequest(rawurl, method string) *HttpRequest {
	var resp http.Response
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Fatal(err)
	}
	req := &http.Request{
		URL:        u,
		Method:     method,
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}

	isLog := atomic.LoadInt32(&OpenConnlog)
	if isLog == 1 {
		trace := &httptrace.ClientTrace{
			GetConn: func(hostPort string) {
				log.Printf("GetConn(%s) need a tcp connection: %s\n", rawurl, hostPort)
			},
			ConnectStart: func(network, addr string) {
				log.Printf("ConnectStart(%s) begin create new connection: %s %s\n", rawurl, network, addr)
			},
			ConnectDone: func(network, addr string, err error) {
				log.Printf("ConnectDone(%s) create connection done: %s %s %v\n", rawurl, network, addr, err)
			},
			GotConn: func(connInfo httptrace.GotConnInfo) {
				log.Printf("GotConn(%s) got a tcp connection: %+v\n", rawurl, connInfo)
			},
			PutIdleConn: func(err error) {
				log.Printf("PutIdleConn(%s) put into idleConn:%v\n", rawurl, err)
			},
		}

		req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	}

	return &HttpRequest{
		url:     rawurl,
		req:     req,
		params:  map[string]string{},
		files:   map[string]string{},
		setting: defaultSetting,
		resp:    &resp,
	}
}

// Get returns *HttpRequest with GET method.
func Get(url string) *HttpRequest {
	return NewRequest(url, "GET")
}

// Post returns *HttpRequest with POST method.
func Post(url string) *HttpRequest {
	return NewRequest(url, "POST")
}

// Put returns *HttpRequest with PUT method.
func Put(url string) *HttpRequest {
	return NewRequest(url, "PUT")
}

// Delete returns *HttpRequest DELETE method.
func Delete(url string) *HttpRequest {
	return NewRequest(url, "DELETE")
}

// Head returns *HttpRequest with HEAD method.
func Head(url string) *HttpRequest {
	return NewRequest(url, "HEAD")
}

// Options returns *HttpRequest with OPTIONS method.
func Options(url string) *HttpRequest {
	return NewRequest(url, "OPTIONS")
}

// HttpSettings
type HttpSettings struct {
	ShowDebug        bool
	UserAgent        string
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
	TlsClientConfig  *tls.Config
	Proxy            func(*http.Request) (*url.URL, error)
	CheckRedirect    func(req *http.Request, via []*http.Request) error
	httpClient       *http.Client
	Transport        http.RoundTripper
	Cookies          http.CookieJar
	EnableCookie     bool
	Gzip             bool
	DumpBody         bool
	ManualSetCookie  bool
}

// HttpRequest provides more useful methods for requesting one url than http.Request.
type HttpRequest struct {
	url     string
	req     *http.Request
	params  map[string]string
	mparams [][2]string
	files   map[string]string
	setting HttpSettings
	resp    *http.Response
	body    []byte
	dump    []byte
}

// get request
func (b *HttpRequest) GetRequest() *http.Request {
	return b.req
}

// get request params
func (b *HttpRequest) GetRequestParams() *map[string]string {
	return &b.params
}

// Change request settings
func (b *HttpRequest) Setting(setting HttpSettings) *HttpRequest {
	b.setting = setting
	return b
}

// func(*http.Request) (*url.URL, error)
func (b *HttpRequest) SetCheckRedirect(selfRedirect func(req *http.Request, via []*http.Request) error) *HttpRequest {
	b.setting.CheckRedirect = selfRedirect
	return b
}

// SetBasicAuth sets the request's Authorization header to use HTTP Basic Authentication with the provided username and password.
func (b *HttpRequest) SetBasicAuth(username, password string) *HttpRequest {
	b.req.SetBasicAuth(username, password)
	return b
}

// SetEnableCookie sets enable/disable cookiejar
func (b *HttpRequest) SetEnableCookie(enable bool) *HttpRequest {
	b.setting.EnableCookie = enable
	return b
}

// SetEnableCookie sets enable/disable cookiejar
func (b *HttpRequest) SetManualSetCookie(enable bool) *HttpRequest {
	b.setting.ManualSetCookie = enable
	return b
}

//Set cookie
func (b *HttpRequest) SetCookieJar(cookiejar http.CookieJar) *HttpRequest {
	b.setting.Cookies = cookiejar
	return b
}

// SetUserAgent sets User-Agent header field
func (b *HttpRequest) SetUserAgent(useragent string) *HttpRequest {
	b.setting.UserAgent = useragent
	return b
}

//Reset request URL, 重置请求的URL
func (b *HttpRequest) ResetRequestURL(rawurl string) *HttpRequest {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Fatal(err)
	}
	b.req.URL = u
	b.url = rawurl
	return b
}

// Reset request Method
func (b *HttpRequest) ResetRequestMethod(method string) *HttpRequest {
	b.req.Method = method
	return b
}

// Debug sets show debug or not when executing request.
func (b *HttpRequest) Debug(isdebug bool) *HttpRequest {
	b.setting.ShowDebug = isdebug
	return b
}

// Dump Body.
func (b *HttpRequest) DumpBody(isdump bool) *HttpRequest {
	b.setting.DumpBody = isdump
	return b
}

// return the DumpRequest
func (b *HttpRequest) DumpRequest() []byte {
	return b.dump
}

// SetTimeout sets connect time out and read-write time out for Request.
func (b *HttpRequest) SetTimeout(connectTimeout, readWriteTimeout time.Duration) *HttpRequest {
	b.setting.ConnectTimeout = connectTimeout
	b.setting.ReadWriteTimeout = readWriteTimeout
	return b
}

// SetTLSClientConfig sets tls connection configurations if visiting https url.
func (b *HttpRequest) SetTLSClientConfig(config *tls.Config) *HttpRequest {
	b.setting.TlsClientConfig = config
	return b
}

//设置自定义的header
func (b *HttpRequest) SetDefaultHeader(defaultHeader http.Header) *HttpRequest {
	b.GetRequest().Header = defaultHeader
	return b
}

func (b *HttpRequest) UncanonicalizedHeader(key, value string) *HttpRequest {
	b.req.Header[key] = []string{value}
	return b
}

func (b *HttpRequest) UncanonicalizedHeaders(headers map[string]string) *HttpRequest {
	for k, v := range headers {
		b.req.Header[k] = []string{v}
	}
	return b
}

// Header add header item string in request.
func (b *HttpRequest) Header(key, value string) *HttpRequest {
	b.req.Header.Set(key, value)
	return b
}

// delete header
func (b *HttpRequest) DelHeader(key string) *HttpRequest {
	b.req.Header.Del(key)
	return b
}

// Headers in request.
func (b *HttpRequest) Headers(headers map[string]string) *HttpRequest {
	for k, v := range headers {
		b.req.Header.Set(k, v)
	}
	return b
}

// Set HOST
func (b *HttpRequest) SetHost(host string) *HttpRequest {
	b.req.Host = host
	return b
}

// Set the protocol version for incoming requests.
// Client requests always use HTTP/1.1.
func (b *HttpRequest) SetProtocolVersion(vers string) *HttpRequest {
	if len(vers) == 0 {
		vers = "HTTP/1.1"
	}

	major, minor, ok := http.ParseHTTPVersion(vers)
	if ok {
		b.req.Proto = vers
		b.req.ProtoMajor = major
		b.req.ProtoMinor = minor
	}

	return b
}

// SetCookie add cookie into request.
func (b *HttpRequest) SetCookie(cookie *http.Cookie) *HttpRequest {
	b.req.Header.Add("Cookie", cookie.String())
	return b
}

// Set transport to
func (b *HttpRequest) SetTransport(transport http.RoundTripper) *HttpRequest {
	b.setting.Transport = transport
	return b
}

// Set http proxy
// example:
//
//	func(req *http.Request) (*url.URL, error) {
// 		u, _ := url.ParseRequestURI("http://127.0.0.1:8118")
// 		return u, nil
// 	}
func (b *HttpRequest) SetProxy(proxy func(*http.Request) (*url.URL, error)) *HttpRequest {
	b.setting.Proxy = proxy
	return b
}

// set AuthProxy
func (b *HttpRequest) SetAuthProxy(proxyUser, proxyPass, proxyIp, ProxyPort string) {
	auth := proxyUser + ":" + proxyPass
	basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	b.Header("Proxy-Authorization", basic)

	proxyUrl := "http://" + proxyIp + ":" + ProxyPort
	b.SetProxy(func(req *http.Request) (*url.URL, error) {
		u, _ := url.ParseRequestURI(proxyUrl)
		u.User = url.UserPassword(proxyUser, proxyPass)
		return u, nil
	})
}

// get http StatusCode
func (b *HttpRequest) GetStatusCode() (int, error) {
	resp, err := b.getResponse()
	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}

// get http location
func (b *HttpRequest) GetLocation() (string, error) {
	resp, err := b.getResponse()
	if err != nil {
		return "", err
	}

	location, err := resp.Location()
	if err != nil {
		return "", err
	}

	return location.String(), nil
}

func (b *HttpRequest) GetResponseParamByName(name string) (string, error) {
	resp, err := b.getResponse()
	if err != nil {
		return "", err
	}
	respParam := resp.Header.Get(name)
	if respParam == "" {
		return "", errors.New("No Find This Response Param!")
	}
	return respParam, nil
}

// Is Forbidden
func (b *HttpRequest) IsForbidden() bool {
	code, err := b.GetStatusCode()
	if err == nil {
		return code == http.StatusForbidden ||
			code == http.StatusRequestTimeout
	} else {
		err_msg := err.Error()
		if strings.Contains(err_msg, "i/o timeout") ||
			strings.Contains(err_msg, "connection refused") {
			return true
		}
	}

	return false
}

// Param adds query param in to request.
// params build query string as ?key1=value1&key2=value2...
func (b *HttpRequest) Param(key, value string) *HttpRequest {
	b.params[key] = value
	return b
}

func (b *HttpRequest) MultiParam(key, value string) *HttpRequest {
	var arg = [2]string{key, value}
	b.mparams = append(b.mparams, arg)
	return b
}

func (b *HttpRequest) PostFile(formname, filename string) *HttpRequest {
	b.files[formname] = filename
	return b
}

// Body adds request raw body.
// it supports string and []byte.
func (b *HttpRequest) Body(data interface{}) *HttpRequest {
	switch t := data.(type) {
	case string:
		bf := bytes.NewBufferString(t)
		b.req.Body = ioutil.NopCloser(bf)
		b.req.ContentLength = int64(len(t))
	case []byte:
		b.req.ContentLength = int64(len(t))
		bf := bytes.NewBuffer(t)
		b.req.Body = ioutil.NopCloser(bf)
	}
	return b
}

// JsonBody adds request raw body encoding by JSON.
func (b *HttpRequest) JsonBody(obj interface{}) (*HttpRequest, error) {
	if b.req.Body == nil && obj != nil {
		buf := bytes.NewBuffer(nil)
		enc := json.NewEncoder(buf)
		if err := enc.Encode(obj); err != nil {
			return b, err
		}
		b.req.Body = ioutil.NopCloser(buf)
		b.req.ContentLength = int64(buf.Len())
		b.req.Header.Set("Content-Type", "application/json")
	}
	return b, nil
}

func (b *HttpRequest) buildUrl(paramBody string) {
	// build GET url with query string
	if b.req.Method == "GET" && len(paramBody) > 0 {
		if strings.Index(b.url, "?") != -1 {
			b.url += "&" + paramBody
		} else {
			b.url = b.url + "?" + paramBody
		}
		return
	}

	// build POST/PUT/PATCH url and body
	if (b.req.Method == "POST" || b.req.Method == "PUT" || b.req.Method == "PATCH") && b.req.Body == nil {
		// with files
		if len(b.files) > 0 {
			pr, pw := io.Pipe()
			bodyWriter := multipart.NewWriter(pw)
			go func() {
				for formname, filename := range b.files {
					fileWriter, err := bodyWriter.CreateFormFile(formname, filename)
					if err != nil {
						log.Fatal(err)
					}
					fh, err := os.Open(filename)
					if err != nil {
						log.Fatal(err)
					}
					//iocopy
					_, err = io.Copy(fileWriter, fh)
					fh.Close()
					if err != nil {
						log.Fatal(err)
					}
				}
				for k, v := range b.params {
					bodyWriter.WriteField(k, v)
				}
				bodyWriter.Close()
				pw.Close()
			}()
			b.Header("Content-Type", bodyWriter.FormDataContentType())
			b.req.Body = ioutil.NopCloser(pr)
			return
		}

		// with params
		if len(paramBody) > 0 {
			b.Header("Content-Type", "application/x-www-form-urlencoded")
			b.Body(paramBody)
		}
	}
}

func (b *HttpRequest) getResponse() (*http.Response, error) {
	if b.resp.StatusCode != 0 {
		return b.resp, nil
	}
	resp, err := b.SendOut()
	if err != nil {
		return nil, err
	}
	b.resp = resp
	if b.setting.ManualSetCookie {
		b.readSetCookies()
	}
	return resp, nil
}

func (b *HttpRequest) SendOut() (*http.Response, error) {
	var paramBody string
	var buf bytes.Buffer

	if len(b.params) > 0 {
		for k, v := range b.params {
			if k != "" {
				buf.WriteString(url.QueryEscape(k))
				buf.WriteByte('=')
			}
			buf.WriteString(url.QueryEscape(v))
			buf.WriteByte('&')
		}
	}
	if len(b.mparams) > 0 {
		for _, p := range b.mparams {
			if len(p) == 2 && p[0] != "" {
				buf.WriteString(url.QueryEscape(p[0]))
				buf.WriteByte('=')
			}
			buf.WriteString(url.QueryEscape(p[1]))
			buf.WriteByte('&')
		}
	}
	if buf.Len() > 0 {
		paramBody = buf.String()
		paramBody = paramBody[0 : len(paramBody)-1]
	}

	b.buildUrl(paramBody)
	tmpURL, err := url.Parse(b.url)
	if err != nil {
		return nil, err
	}

	b.req.URL = tmpURL

	if b.setting.ManualSetCookie {
		b.addCooike()
	}

	trans := b.setting.Transport

	if trans == nil {
		// create default transport
		trans = &http.Transport{
			TLSClientConfig: b.setting.TlsClientConfig,
			Proxy:           b.setting.Proxy,
			//Dial:            TimeoutDialer(b.setting.ConnectTimeout, b.setting.ReadWriteTimeout),
			DialContext: (&net.Dialer{
				Timeout:   b.setting.ConnectTimeout, //  30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
		}
	} else {
		// if b.transport is *http.Transport then set the settings.
		if t, ok := trans.(*http.Transport); ok {
			t.TLSClientConfig = b.setting.TlsClientConfig
			t.Proxy = b.setting.Proxy
			t.DialContext = (&net.Dialer{
				Timeout:   b.setting.ConnectTimeout, //  30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext

		}
	}

	var jar http.CookieJar = nil
	if b.setting.EnableCookie && !b.setting.ManualSetCookie {
		if b.setting.Cookies == nil {
			b.setting.Cookies, _ = cookiejar.New(nil)
		}
		jar = b.setting.Cookies
	}

	client := &http.Client{
		Transport:     trans,
		Jar:           jar,
		CheckRedirect: b.setting.CheckRedirect,
		Timeout:       b.setting.ReadWriteTimeout,
	}

	if b.setting.UserAgent != "" {
		if _, ok := b.req.Header["User-Agent"]; !ok {
			b.req.Header.Set("User-Agent", b.setting.UserAgent)
		}
	}

	if b.req.Header.Get("Host") != "" {
		b.req.Host = b.req.Header.Get("Host")
	}

	if b.setting.ShowDebug {
		dump, err := httputil.DumpRequest(b.req, b.setting.DumpBody)
		if err != nil {
			log.Println(err.Error())
		}
		b.dump = dump
	}
	return client.Do(b.req)
}

// 获取请求body 附带检查code
func (b *HttpRequest) StringWithCheckCode() (string, error) {
	data, err := b.Bytes()
	if err != nil {
		return "", err
	}

	code, err := b.GetStatusCode()
	if err != nil {
		return string(data), err
	}

	if code >= http.StatusBadRequest {
		return string(data), fmt.Errorf("response error:%d %s", code, http.StatusText(code))
	}

	return string(data), nil
}

// String returns the body string in response.
// it calls Response inner.
func (b *HttpRequest) String() (string, error) {
	data, err := b.Bytes()
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Bytes returns the body []byte in response.
// it calls Response inner.
func (b *HttpRequest) Bytes() ([]byte, error) {
	if b.body != nil {
		return b.body, nil
	}
	resp, err := b.getResponse()
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	if b.setting.Gzip && resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		b.body, err = ioutil.ReadAll(reader)
	} else {
		b.body, err = ioutil.ReadAll(resp.Body)
	}
	return b.body, err
}

// ToFile saves the body data in response to one file.
// it calls Response inner.
func (b *HttpRequest) ToFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	resp, err := b.getResponse()
	if err != nil {
		return err
	}
	if resp.Body == nil {
		return nil
	}
	defer resp.Body.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}

// ToJson returns the map that marshals from the body bytes as json in response .
// it calls Response inner.
func (b *HttpRequest) ToJson(v interface{}) error {
	data, err := b.Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// ToXml returns the map that marshals from the body bytes as xml in response .
// it calls Response inner.
func (b *HttpRequest) ToXml(v interface{}) error {
	data, err := b.Bytes()
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, v)
}

// Response executes request client gets response mannually.
func (b *HttpRequest) Response() (*http.Response, error) {
	return b.getResponse()
}

// Reset reset HttpRequest to its initial state
func (b *HttpRequest) Reset() {
	var resp http.Response
	b.resp = &resp
	b.body = nil
	b.dump = nil
}

// TimeoutDialer returns functions of connection dialer with timeout settings for http.Transport Dial field.
func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		err = conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, err
	}
}
