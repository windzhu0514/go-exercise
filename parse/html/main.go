package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
)

func main() {
	data, err := ioutil.ReadFile("./parse/html/test.html")
	if err != nil {
		panic(err)
	}

	//fmt.Println(parseHTML(string(data)))
	//parseScript(string(data))
	fmt.Println(parseorderdetails(string(data)))
}

func parseHTML(html string) error {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil
	}

	params := make(map[string]string)
	doc.Find("body form input").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		value, _ := s.Attr("value")
		if name != "" && value != "" {
			params[name] = value
		}
	})

	fmt.Println(params)

	return nil
}

type COPassengerInfo struct {
	Name     string `json:"name"`               // 乘客姓名
	Type     int    `json:"type"`               // 乘客类型 成人：ADULT 儿童：CHILD 婴儿：INFANT
	Birthday string `json:"birthday"`           // 出生日期 "yyyy-MM-dd"
	CertNo   string `json:"certNo"`             // 证件号码
	CertType int    `json:"certType,omitempty"` // 证件类型  NI:身份证 JG:军官证 ID:其它 PP:护照
	Mobile   string `json:"mobile,omitempty"`   // 号码
	FareInfo
}

type FareInfo struct {
	Price      string `json:"price"`      // 价格
	AirportTax string `json:"airportTax"` // 机场建设费
	FuelTax    string `json:"fuelTax"`    // 燃油税
	OtherTax   string `json:"otherTax"`   // 其他费用
}

type Ticket struct {
	COPassengerInfo
	SegmentList []Ticket_Segment `json:"segmentList"`
}

type Ticket_Segment struct {
	RouteNo         int    `json:"routeNo"`                 // 航段编号 1代表第一段2代表第二段 依次类推
	Departure       string `json:"departure"`               // 出发城市 "上海"
	DepartureCode   string `json:"departureCode,omitempty"` // 出发城市三字码 "SHA"
	DepAirportName  string `json:"depAirportName"`          // 出发机场 "虹桥国际机场T1"
	DepAirport      string `json:"depAirport"`              // 出发机场三字码 "SHA"
	Destination     string `json:"destination"`             // 到达城市 "深圳"
	DestinationCode string `json:"destinationCode"`         // 到达城市三字码 "SZX"
	DestAirportName string `json:"destAirportName"`         // 到达机场 "宝安国际机场T3"
	DestAirport     string `json:"destAirport"`             // 到达机场三字码 "SZX"
	FlightNo        string `json:"flightNo"`                // 航班号 "9C8917"
	FlightType      string `json:"flightType"`              // 航班机型 "空客320_186"
	CabinName       string `json:"cabinName"`               // 舱位名 "R3E" 会有重复
	CabinCode       string `json:"cabinCode"`               // 舱位代码 "5"
	DepartureTime   string `json:"departureTime"`           // 起飞时间 "2020-06-29 06:30:00"
	ArrivalTime     string `json:"arrivalTime"`             // 到达时间 "2020-06-29 08:55:00"
	MainFlightNo    string `json:"MainFlightNo"`            // 承载航班号
}

type orderDetail struct {
	OrderCreateTime string
	ContactName     string
	ContactMobile   string
	ContactEmail    string
	TicketList      []Ticket
}

func parseorderdetails(html string) (od orderDetail, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return od, err
	}

	od.OrderCreateTime = doc.Find("body div.p-wrap.f-cb div div.p-holder div.p-content div.p-order" +
		"div.add-erweima.clearfix table.m-order tr:nth-child(4) td:nth-child(2)").Text()
	od.OrderCreateTime = strings.TrimSpace(od.OrderCreateTime)

	sectionOrder := doc.Find("body div.p-wrap.f-cb div div.p-holder div.p-content div.p-order")

	var departure, departureTime, depAirportName, destination, arrivalTime, destAirportName string
	departureInfo := sectionOrder.Find("table.m-route.c-table tbody tr.last td:nth-child(2) ul li:nth-child(1) p").Text()
	fields := strings.Fields(departureInfo)
	if len(fields) > 3 {
		departureTime = fields[1] + ":00"
		departure = fields[2]
		depAirportName = fields[3]
	}
	destinationInfo := sectionOrder.Find("table.m-route.c-table tbody tr.last td:nth-child(2) ul li:nth-child(2) p").Text()
	fields = strings.Fields(destinationInfo)
	if len(fields) > 3 {
		arrivalTime = fields[1] + ":00"
		destination = fields[2]
		destAirportName = fields[3]
	}

	sectionService := doc.Find("body div.p-wrap.f-cb div.p-holder div.p-content div.p-order table.m-service.c-table " +
		"tbody tr:nth-child(2)")
	flightNo := strings.TrimSpace(sectionService.Find("td:nth-child(1)").Text())
	sectionService.Find("td:nth-child(2) h6").Each(func(i int, s *goquery.Selection) {
		var ticket common.Ticket
		var passager common.COPassengerInfo
		var segment common.Ticket_Segment
		sectonCService := sectionService.Find("div.c-service").Eq(i)
		segment.FlightNo = flightNo
		segment.CabinName = sectonCService.Find("ul > li:nth-child(1) > div > span.c-name > span").Text()
		segment.CabinName = strings.TrimSpace(segment.CabinName)
		segment.CabinName = strings.Trim(segment.CabinName, "()")
		segment.Departure = departure
		segment.DepartureTime = departureTime
		segment.DepAirportName = depAirportName
		segment.Destination = destination
		segment.ArrivalTime = arrivalTime
		segment.DestAirportName = destAirportName

		passager.Name = strings.TrimPrefix(strings.TrimSpace(s.Text()), "乘机人：")
		passager.Price = strings.TrimSpace(sectonCService.Find("ul > li:nth-child(1) > div > span.c-price").Text())
		passager.Price = strings.TrimPrefix(passager.Price, "/¥")
		passager.AirportTax = strings.TrimSpace(sectonCService.Find("ul > li:nth-child(4) > div > span.c-price").Text())
		passager.AirportTax = strings.TrimPrefix(passager.AirportTax, "/¥")
		passager.FuelTax = strings.TrimSpace(sectonCService.Find("ul > li:nth-child(6) > div > span.c-price").Text())
		passager.FuelTax = strings.TrimPrefix(passager.FuelTax, "/¥")
		passager.OtherTax = strings.TrimSpace(sectonCService.Find("ul > li:nth-child(8) > div > span.c-price").Text())
		passager.OtherTax = strings.TrimPrefix(passager.OtherTax, "/¥")

		sectionPassengers := doc.Find("body div.p-wrap.f-cb div.p-content div.p-passenger").Eq(0)
		sectionPassengers.Find("ul li").EachWithBreak(func(i int, ss *goquery.Selection) bool {
			name := ss.Find("table tr td.name.f-cb p:nth-child(2)").Text()
			name = strings.TrimSpace(name)
			if name == passager.Name {
				pType := ss.Find("table tr td.name.f-cb p.c-ps").Text()
				pType = strings.TrimSpace(pType)
				switch pType {
				case "(成人)":
					passager.Type = common.PassengerTypeAdult
				case "(儿童)":
					passager.Type = common.PassengerTypeChild
				case "(婴儿)":
					passager.Type = common.PassengerTypeInfant
				}

				passager.Birthday = ss.Find("table tr td.info p").Eq(0).Text()
				passager.Birthday = strings.TrimPrefix(strings.TrimSpace(passager.Birthday), "出生年月：")
				passager.Mobile = ss.Find("table tr td.info div.show_Tips.pMobile p span").Text()
				passager.Mobile = strings.TrimSpace(passager.Mobile)
				passager.CertNo = ss.Find("table tr td.info div.show_Tips.certificateTypeV p span.certificate.hide-certificate").Text()
				passager.CertNo = strings.TrimSpace(passager.CertNo)
				certType := ss.Find("table tr td.info div.show_Tips.certificateTypeV p span.hide-certificate-type").Text()
				certType = strings.TrimSpace(certType)
				passager.CertType = formatCertType(certType)
				return false
			}

			return true
		})

		ticket.COPassengerInfo = passager
		ticket.SegmentList = append(ticket.SegmentList, segment)
		od.TicketList = append(od.TicketList, ticket)
	})

	od.ContactName = doc.Find("body div.p-wrap.f-cb div div.p-holder div.p-content div:nth-child(3) ul li table tbody tr td.name.f-cb p").Text()
	od.ContactName = strings.TrimSpace(od.ContactName)
	od.ContactMobile = strings.TrimSpace(doc.Find("#concactPerson").Text())

	return od, nil
}

func parseScript(html string) error {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return err
	}

	script := doc.Find("body script").Text()
	indexb := strings.Index(script, "{")
	indexe := strings.Index(script, "document.cookie = \"")
	if indexb < 0 || indexe < 0 {
		return errors.New("无法解析脚本")
	}
	indexee := strings.Index(script[indexe+19:], "=\"")
	if indexee < 0 {
		return errors.New("无法解析脚本")
	}

	cookieName := script[indexe+19 : indexe+19+indexee]
	fmt.Println(cookieName)

	script = script[indexb+1 : indexe]
	fmt.Println(strings.TrimSpace(script))
	vm := otto.New()
	if _, err = vm.Run(script); err != nil {
		return err
	}

	v, err := vm.Get("v")
	if err != nil {
		return err
	}

	fmt.Println(v.String())
	return nil
}
