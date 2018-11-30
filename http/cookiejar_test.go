package http

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"
)

func TestCookieJar(t *testing.T) {
	jar, _ := cookiejar.New(nil)
	setCookie(jar, "http://www.baidu.com", "aaa", "1111")
	printCookie(jar, "http://www.baidu.com")

	setCookie(jar, "http://www.xxx.baidu.com", "aaa", "1111")
	printCookie(jar, "http://www.xxx.baidu.com")

	setCookie(jar, "http://www.yyy.zzz.baidu.com", "aaa", "1111")
	printCookie(jar, "http://www.yyy.zzz.baidu.com")
}

func setCookie(jar *cookiejar.Jar, rawurl, name, value string) {
	if jar == nil {
		return
	}

	URL, err := url.Parse(rawurl)
	if err != nil {
		return
	}

	cookies := jar.Cookies(URL)
	for i, v := range cookies {
		if v.Name == name {
			cookies[i].Value = value

			jar.SetCookies(URL, cookies)

			return
		}
	}

	cookies = append(cookies, &http.Cookie{Name: name, Value: value})

	jar.SetCookies(URL, cookies)
}

func printCookie(jar *cookiejar.Jar, rawurl string) {
	URL, err := url.Parse(rawurl)
	if err != nil {
		return
	}

	cookies := jar.Cookies(URL)
	fmt.Println(cookies)
}

func getCookie(jar *cookiejar.Jar, rawurl, name string) string {
	URL, err := url.Parse(rawurl)
	if err != nil {
		return ""
	}

	cookies := jar.Cookies(URL)
	for _, v := range cookies {
		if v.Name == name {
			return v.Value
		}
	}

	return ""
}
