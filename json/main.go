package main

import (
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
	Departure string
	// json字段名默认为导出的结构体字段名
	Destination string `json:"destination"`
	//DepartureCode string
}

type omit *struct{}

func main() {
	var busInfo = QueryCoaches_BusInfo{}

	busInfo.Departure = "苏州"
	busInfo.Destination = "河南"
	//busInfo.DepartureCode = "1231231"
	// busInfo.Destination = "三岔"
	// busInfo.ScheduleId = 2999494
	// busInfo.DepartureCode = 1129944
	// busInfo.TicketPrice = 12.5
	// busInfo.TicketPrice2 = 15.6
	//

	jsonData, err := json.Marshal(struct {
		Info        QueryCoaches_BusInfo
		Destination string `json:"destination,omitempty"`
	}{Info: busInfo})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))

	// var buff bytes.Buffer
	// enc := json.NewEncoder(&buff)
	// enc.SetEscapeHTML(false)
	// err := enc.Encode(busInfo)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// fmt.Println(buff.String())
	// jsonstr := `{"departure":"三门","scheduleId":2999494,"departureCode":1129944,"ticketPrice":"12.5"}`
	// fmt.Println(json.Unmarshal([]byte(jsonstr), &busInfo))
	// fmt.Println(busInfo)
}
