package main

type structName struct {
	FulfillChangeRulesFlightDTOs []struct {
		DestAirportCode  string `json:"destAirportCode"`
		DestAirportCode1 string `json:"destAirportCode1"`
		DestAirportCode2 string `json:"destAirportCode2"`
		DestAirportName  string `json:"destAirportName"`
		DestAirportName1 string `json:"destAirportName1"`
		DestAirportName2 string `json:"destAirportName2"`
		DestCityCode     string `json:"destCityCode"`
		DestCityCode1    string `json:"destCityCode1"`
		DestCityCode2    string `json:"destCityCode2"`
		DestCityName     string `json:"destCityName"`
		DestCityName1    string `json:"destCityName1"`
		DestCityName2    string `json:"destCityName2"`
		DestTimeBJ       string `json:"destTimeBJ"`
		DestTimeLocal    string `json:"destTimeLocal"`
		DestTimeLocalSys int64  `json:"destTimeLocalSys"`
		FlightNo         string `json:"flightNo"`
		OriAirportCode   string `json:"oriAirportCode"`
		OriAirportName   string `json:"oriAirportName"`
		OriCityCode      string `json:"oriCityCode"`
		OriCityName      string `json:"oriCityName"`
		OriTimeBJ        string `json:"oriTimeBJ"`
		OriTimeLocal     string `json:"oriTimeLocal"`
		OriTimeLocalSys  int64  `json:"oriTimeLocalSys"`
		SegHeadID        int64  `json:"segHeadId"`
	} `json:"fulfillChangeRulesFlightDTOs"`
	StatusCode string `json:"statusCode"`
}

func test() {
	var data structName
	data.FulfillChangeRulesFlightDTOs = nil
	data.StatusCode = ""
}
