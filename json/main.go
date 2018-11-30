package main

import (
	"encoding/json"
	"fmt"
)

type QueryCoaches_BusInfo struct {
	DepartureCode          string `json:"departureCode"`                 // 出发城市编码
	DestinationCode        string `json:"destinationCode"`               // 到达城市编码
	ScheduleId             string `json:"scheduleId,omitempty"`          // 班次唯一id
	CoachType              string `json:"coachType"`                     // 车辆类型
	CoachNo                string `json:"coachNo"`                       // 班次编码(车次号)
	Seattype               string `json:"seatType"`                      // 座位类型
	Departure              string `json:"departure"`                     // 出发城市
	Destination            string `json:"destination"`                   // 到达城市
	DptStation             string `json:"dptStation"`                    // 出发车站
	ArrStation             string `json:"arrStation"`                    // 到达车站
	Terminal               string `json:"terminal"`                      // 终点城市
	TerminalStation        string `json:"terminalStation"`               // 终点车站
	TicketPrice            string `json:"ticketPrice"`                   // 全票票价
	TicketLeft             string `json:"ticketLeft"`                    // 余票数量
	CanBooking             bool   `json:"canBooking"`                    // 是否可以订票
	DptStationCode         string `json:"dptStationCode"`                // 出发车站编码
	DestinationStationCode string `json:"destinationStationCode"`        // 到达车站编码
	Distance               string `json:"distance"`                      // 行车里程
	IsHomeDelivery         bool   `json:"isHomeDelivery"validate:"true"` // 是否始发站
	IsPassingStation       bool   `json:"isPassingStation"`              // 是否途径站
	DptDate                string `json:"dptDate"`                       // 出发日期
	DptTime                string `json:"dptTime"`                       // 出发时间
	RunTime                string `json:"runTime"`                       // 行车时间
}

func main() {
	jsonStr := `[{"departureCode":"","destinationCode":"","coachType":"中中一","coachNo":"1712","seatType":"","departure":"三门","destination":"三岔","dptStation":"三门客运中心","arrStation":"三岔","terminal":"","terminalStation":"蛇蟠普通","ticketPrice":"6","ticketLeft":"17","canBooking":true,"dptStationCode":"100201","destinationStationCode":"三岔","distance":"","isHomeDelivery":true,"isPassingStation":true,"dptDate":"2018-10-20","dptTime":"12:00","runTime":""},{"departureCode":"","destinationCode":"","coachType":"中中一","coachNo":"1712","seatType":"","departure":"三门","destination":"三岔","dptStation":"三门客运中心","arrStation":"三岔","terminal":"","terminalStation":"蛇蟠普通","ticketPrice":"6","ticketLeft":"17","canBooking":true,"dptStationCode":"100201","destinationStationCode":"三岔","distance":"","isHomeDelivery":true,"isPassingStation":true,"dptDate":"2018-10-20","dptTime":"12:00","runTime":""}]`

	var GSBus []QueryCoaches_BusInfo
	json.Unmarshal([]byte(jsonStr), &GSBus)

	jsonData, err := json.Marshal(GSBus[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

}
