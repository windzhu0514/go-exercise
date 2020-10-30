package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	data, err := ioutil.ReadFile("./parse/html/test.html")
	if err != nil {
		panic(err)
	}

	//fmt.Println(parseHTML(string(data)))
	//parseScript(string(data))
	fmt.Println(parseHTML(string(data)))
	//fmt.Println(SplitSeatNo(""))
	//fmt.Println(SplitSeatNo("H"))
}

func SplitSeatNo(seatNo string) (row, col string) {
	row = seatNo[0 : len(seatNo)-1]
	col = seatNo[len(seatNo)-1:]
	return
}

func parseHTML(html string) error {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	errMsg := doc.Find("body div.container div.message").Text()
	errMsg = strings.ReplaceAll(errMsg, "\n", "")
	errMsg = strings.ReplaceAll(errMsg, " ", "")
	fmt.Println(errMsg)

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
