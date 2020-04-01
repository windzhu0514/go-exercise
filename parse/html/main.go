package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
)

func main() {
	data, err := ioutil.ReadFile("./test.html")
	if err != nil {
		panic(err)
	}

	fmt.Println(parseBookInfo(string(data)))
	//parseScript(string(data))
}

type CheckTicketRes_Data struct {
	TicketNo         string              `json:"ticketNo"`         // 电子客票号
	BookingReference string              `json:"bookingReference"` // PNR
	FullName         string              `json:"fullName"`         // 乘客全名(英文含/,中文无/)
	FirstName        string              `json:"firstName"`        // 乘客名
	LastName         string              `json:"lastName"`         // 乘客姓
	Gender           string              `json:"gender"`           // 乘客姓别 M：男；F：女
	AdultTicketPrice Monetary            `json:"adultTicketPrice"` // 成人票价
	AdultTotalTax    Monetary            `json:"adultTotalTax"`    // 成人税额
	TaxDetail        map[string]Monetary `json:"taxDetail"`        // 税详细，不同航司key不同
	Segments         []FlightSegment     `json:"segments"`         // 航段集合
}

type CheckTicketPassengerInfo struct {
	FirstName string `json:"firstName"` // 乘客名
	LastName  string `json:"lastName"`  // 乘客姓
	CardType  string `json:"cardType"`  // 证件号码
	CardNo    string `json:"cardNo"`    // 证件号码
}

type Monetary struct {
	Currency    string `json:"currency"`    // 币种
	TotalAmount string `json:"totalAmount"` // 总金额
}

type FlightSegment struct {
	Index       string `json:"index"`       // 从1开始
	FlightNo    string `json:"flightNo"`    // 航班号
	CabinCode   string `json:"cabinCode"`   // 舱位
	DepAirport  string `json:"depAirport"`  // 出发机场三字码
	ArrAirport  string `json:"arrAirport"`  // 到达机场三字码
	DepDateTime string `json:"depDateTime"` // 起飞时间	yyyy-MM-dd HH:mm:ss
	ArrDateTime string `json:"arrDateTime"` // 到达时间 yyyy-MM-dd HH:mm:ss
	// 航段状态 UNKNOWN:未知 CANCELED:已取消 SOLD:已售票 REFUND:申请退票 USED:已乘机
	Status string `json:"status"`
}

func HasHan(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

func IsAllASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= utf8.RuneSelf {
			return false
		}
	}
	return true
}

func parseBookInfo(html string) (data CheckTicketRes_Data, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return data, err
	}

	bFind := false
	bIsChild := false
	coutAdult := 0
	sections := doc.Find("form#ManageBookingForm table.print-guest-details tbody tr.guest-padding")
	for i := 0; i < sections.Length(); i += 2 {
		guestName := sections.Eq(i).Find("td").Text()
		fields := strings.Fields(guestName)
		if len(fields) < 5 {
			return data, errors.New("姓名格式无法解析")
		}

		adultorchild := fields[4]
		if strings.Contains(strings.ToLower(adultorchild), "adult") {
			coutAdult++
		}

		gender := fields[1]
		if strings.Contains(strings.ToLower(gender), "ms") {
			data.Gender = "F"
		} else {
			data.Gender = "M"
		}

		firstName := fields[2]
		lastName := fields[3]
		if firstName != "YAN" &&
			lastName != "LI" {
			continue
		}

		bFind = true

		data.FirstName = firstName
		data.LastName = lastName
		if IsAllASCII(firstName) && IsAllASCII(lastName) {
			data.FullName = lastName + "/" + firstName
		} else {
			data.FullName = lastName + firstName
		}

		if strings.Contains(strings.ToLower(adultorchild), "child") {
			bIsChild = true
		}
	}

	if !bFind {
		return data, errors.New("未找到指定的姓名")
	}

	var segment FlightSegment
	segment.Index = "1"

	fightInfo := doc.Find("form#ManageBookingForm table.print-flight-details tbody tr td")
	segment.FlightNo = fightInfo.Eq(0).Find("p").Eq(1).Text()
	segment.FlightNo = strings.TrimSpace(segment.FlightNo)
	segment.FlightNo = strings.ReplaceAll(segment.FlightNo, " ", "")

	depFightText := fightInfo.Eq(1).Text()
	fields := strings.FieldsFunc(depFightText, func(r rune) bool {
		if r == rune('\r') || r == rune('\n') {
			return true
		}
		return false
	})
	if len(fields) >= 3 {
		cityAndCode := strings.Split(fields[0], "|")
		if len(cityAndCode) >= 2 {
			segment.DepAirport = strings.TrimSpace(cityAndCode[1])
		}

		//  Tue. 14 Jan. 2020, 0320H (03:20AM)
		depDateTime := strings.TrimSpace(fields[2])
		indexComma := strings.Index(depDateTime, ", ")
		if indexComma < 0 {
			return data, errors.New("无法解析出发时间：" + depDateTime)
		}

		indexSpace := strings.LastIndex(depDateTime, " ")
		if indexSpace < 0 || indexSpace < indexComma {
			return data, errors.New("无法解析出发时间：" + depDateTime)
		}

		depDateTime = depDateTime[:indexComma] + depDateTime[indexSpace:]
		dateTime, err := time.Parse("Mon. 02 Jan. 2006 (15:04AM)", depDateTime)
		if err != nil {
			return data, errors.New("无法解析出发时间：" + depDateTime)
		}

		segment.DepDateTime = dateTime.Format("2006-01-02 15:04:00")
	}

	arrFightText := fightInfo.Eq(3).Text()
	fields = strings.FieldsFunc(arrFightText, func(r rune) bool {
		if r == rune('\r') || r == rune('\n') {
			return true
		}
		return false
	})
	if len(fields) >= 3 {
		cityAndCode := strings.Split(fields[0], "|")
		if len(cityAndCode) >= 2 {
			segment.ArrAirport = strings.TrimSpace(cityAndCode[1])
		}

		// Tue. 14 Jan. 2020, 0545H (05:45AM)
		arrDateTime := strings.TrimSpace(fields[2])

		indexComma := strings.Index(arrDateTime, ", ")
		if indexComma < 0 {
			return data, errors.New("无法解析到达时间：" + arrDateTime)
		}

		indexSpace := strings.LastIndex(arrDateTime, " ")
		if indexSpace < 0 || indexSpace < indexComma {
			return data, errors.New("无法解析到达时间：" + arrDateTime)
		}

		arrDateTime = arrDateTime[:indexComma] + arrDateTime[indexSpace:]
		dateTime, err := time.Parse("Mon. 02 Jan. 2006 (15:04AM)", arrDateTime)
		if err != nil {
			return data, errors.New("无法解析到达时间：" + arrDateTime)
		}

		segment.ArrDateTime = dateTime.Format("2006-01-02 15:04:00")
	}

	status := doc.Find("form#ManageBookingForm div.print-booking-details div").Eq(0).Find("strong").Text()
	segment.Status = MapFlightStatus(strings.TrimSpace(status))

	data.Segments = append(data.Segments, segment)

	pnr := doc.Find("form#ManageBookingForm div.print-booking-details div").Eq(2).Find("strong").Text()
	data.BookingReference = strings.TrimSpace(pnr)

	if bIsChild {
		return data, nil
	}

	// 费用
	paymentInfo := doc.Find("form#ManageBookingForm table.print-payment-details tbody tr").
		Eq(0).Find("td").Eq(3).Text()
	paymentInfo = strings.TrimSpace(paymentInfo)
	fields = strings.Fields(paymentInfo)
	if len(fields) >= 2 {
		currency := fields[0]
		totalAmountStr := strings.Replace(fields[1], ",", "", -1)
		data.AdultTicketPrice.Currency = currency
		totalAmount, err := strconv.ParseFloat(totalAmountStr, 64)
		if err != nil {
			return data, errors.New("")
		}
		data.AdultTicketPrice.TotalAmount = strconv.FormatFloat(totalAmount/float64(coutAdult), 'f', 2, 64)
	}

	return data, nil
}

func MapFlightStatus(s string) string {
	switch s {
	case "CONFIRMED":
		return "SOLD" // 已售票
	case "":
		return "USED" // 已乘机
	case " ":
		return "CANCELED" // 已取消
	case "  ":
		return "REFUND" // 申请退票
	case "   ":
		return "UNKNOWN" // 未知
	}

	return ""
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
