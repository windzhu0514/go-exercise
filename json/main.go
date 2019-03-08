package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// 编码后如果出现了相同的json key，这个key不会编码到json字符串中。
// 结果 {}
type price struct {
	TicketPrice  float64 `json:"ticketPrice"`
	TicketPrice2 float32 `json:"ticketPrice"`
}

type QueryCoaches_BusInfo struct {
	Departure string `json:"departure"`
	// json字段名默认为导出的结构体字段名
	Destination   string `json:"-,"`
	DptStation    string `json:"-,"`
	ScheduleId    int    `json:"scheduleId,omitempty"`
	DepartureCode string `json:"departureCode"`
	// string选项表示在json编码时，字段编码为JSON字符串。只支持类型是string,浮点型,所有整形,bool
	TicketPrice  float64 `json:"ticketPrice,string"`
	TicketPrice2 float32 `json:"ticketPrice2"`
	CanBook      bool    `json:"ticketPrice2"`
}

func main() {
	var busInfo QueryCoaches_BusInfo
	busInfo.Departure = "sdjkfj@#$#@$CC>LK:^&%$#$SFSFSFSF$#@$#@$"
	// busInfo.Destination = "三岔"
	// busInfo.ScheduleId = 2999494
	// busInfo.DepartureCode = 1129944
	// busInfo.TicketPrice = 12.5
	// busInfo.TicketPrice2 = 15.6
	//
	// jsonData, err := json.Marshal(busInfo)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//fmt.Println(string(jsonData))
	var buff bytes.Buffer
	enc := json.NewEncoder(&buff)
	enc.SetEscapeHTML(false)
	err := enc.Encode(busInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(buff.String())
	// jsonstr := `{"departure":"三门","scheduleId":2999494,"departureCode":1129944,"ticketPrice":"12.5"}`
	// fmt.Println(json.Unmarshal([]byte(jsonstr), &busInfo))
	// fmt.Println(busInfo)
}

func JsonMarshalNoError(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(data)
}
