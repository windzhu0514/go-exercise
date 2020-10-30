package xml

import (
	"encoding/xml"
	"fmt"
	"testing"
)

// <?xml version='1.0' encoding='UTF-8' standalone='yes' ?>
//<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
//  <soap:Header MOBILE_TYPE="hxZghpJHDZHlPbaKiX0birx2LDBnIlIzAuK0L1gRbnVysJFT6DilqY0R9eSkYPYX7SbPIVAV4T8dZpPG7VhTNQ==" />
//  <soap:Body>
//    <ns2:queryTrip xmlns:ns2="http://com/shenzhenair/mobilewebservice/booking">
//      <QUERY_TRIP_CONDITION>
//        <CERT_NO>4795466349867</CERT_NO>
//        <MOBILE_NO>13506209670</MOBILE_NO>
//        <PASSENGER_NAME>曾涛</PASSENGER_NAME>
//      </QUERY_TRIP_CONDITION>
//    </ns2:queryTrip>
//  </soap:Body>
//</soap:Envelope>

type queryTripData struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"xmlns:soap,attr"`
	Header  struct {
		Text       string `xml:",chardata"`
		MOBILETYPE string `xml:"MOBILE_TYPE,attr"`
	} `xml:"soap:Header"`
	Body struct {
		Text      string `xml:",chardata"`
		QueryTrip struct {
			Text               string `xml:",chardata"`
			Ns2                string `xml:"xmlns:ns2,attr"`
			QUERYTRIPCONDITION struct {
				Text          string `xml:",chardata"`
				CERTNO        string `xml:"CERT_NO"`
				MOBILENO      string `xml:"MOBILE_NO"`
				PASSENGERNAME string `xml:"PASSENGER_NAME"`
			} `xml:"QUERY_TRIP_CONDITION"`
		} `xml:"ns2:queryTrip"`
	} `xml:"soap:Body"`
}

func TestEncode(t *testing.T) {
	var e queryTripData
	e.Soap = "http://schemas.xmlsoap.org/soap/envelope/"
	e.Header.MOBILETYPE = "hxZghpJHDZHlPbaKiX0birx2LDBnIlIzAuK0L1gRbnVysJFT6DilqY0R9eSkYPYX7SbPIVAV4T8dZpPG7VhTNQ=="
	e.Body.QueryTrip.Ns2 = "http://com/shenzhenair/mobilewebservice/booking"
	e.Body.QueryTrip.QUERYTRIPCONDITION.PASSENGERNAME = "曾涛"
	e.Body.QueryTrip.QUERYTRIPCONDITION.CERTNO = "4795466349867"
	e.Body.QueryTrip.QUERYTRIPCONDITION.MOBILENO = "13506209670"

	data, err := xml.MarshalIndent(e, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(data))
	//	xmlStr := `<?xml version='1.0' encoding='UTF-8' standalone='yes' ?>
	//<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	//  <soap:Header MOBILE_TYPE="hxZghpJHDZHlPbaKiX0birx2LDBnIlIzAuK0L1gRbnVysJFT6DilqY0R9eSkYPYX7SbPIVAV4T8dZpPG7VhTNQ==" />
	//  <soap:Body>
	//    <ns2:queryTrip xmlns:ns2="http://com/shenzhenair/mobilewebservice/booking">
	//      <QUERY_TRIP_CONDITION>
	//        <CERT_NO>4795466349867</CERT_NO>
	//        <MOBILE_NO>13506209670</MOBILE_NO>
	//        <PASSENGER_NAME>曾涛</PASSENGER_NAME>
	//      </QUERY_TRIP_CONDITION>
	//    </ns2:queryTrip>
	//  </soap:Body>
	//</soap:Envelope>`
	//	var e Envelope
	//	if err := xml.Unmarshal([]byte(xmlStr), &e); err != nil {
	//		t.Fatal(err)
	//	}
	//	fmt.Printf("%+v\n", e)
	//	data, err := xml.Marshal(e)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	fmt.Println(xml.Header + string(data))
}

type queryTripResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text              string `xml:",chardata"`
		QueryTripResponse struct {
			Text            string `xml:",chardata"`
			Ns2             string `xml:"ns2,attr"`
			QUERYTRIPRESULT struct {
				Text     string `xml:",chardata"`
				OPRESULT string `xml:"OP_RESULT"`
				FLIGHT   struct {
					Text              string `xml:",chardata"`
					AIRLINECODE       string `xml:"AIRLINE_CODE"`
					ARRIVEDATE        string `xml:"ARRIVE_DATE"`
					ARRIVETIME        string `xml:"ARRIVE_TIME"`
					BIRTHDAY          string `xml:"BIRTH_DAY"`
					CANCELCHECKINFLAG string `xml:"CANCEL_CHECK_IN_FLAG"`
					CARRAIRLINECODE   string `xml:"CARR_AIRLINE_CODE"`
					CERTNO            string `xml:"CERT_NO"`
					CERTTYPE          string `xml:"CERT_TYPE"`
					CHECKINFLAG       string `xml:"CHECK_IN_FLAG"`
					CHECKINPLAT       string `xml:"CHECK_IN_PLAT"`
					CLASSCODE         string `xml:"CLASS_CODE"`
					CLASSLEVEL        string `xml:"CLASS_LEVEL"`
					ISAMERICAN        string `xml:"IS_AMERICAN"`
					DSTTERM           string `xml:"DST_TERM"`
					EDICKI            string `xml:"EDI_CKI"`
					FLIGHTDATE        string `xml:"FLIGHT_DATE"`
					FLIGHTNO          string `xml:"FLIGHT_NO"`
					FLIGHTNUMBER      string `xml:"FLIGHT_NUMBER"`
					FLIGHTTIME        string `xml:"FLIGHT_TIME"`
					FROMCITY          string `xml:"FROM_CITY"`
					FROMCITYAIRPORT   string `xml:"FROM_CITY_AIRPORT"`
					ISOPENCITY        string `xml:"IS_OPEN_CITY"`
					FROMCITYNAME      string `xml:"FROM_CITY_NAME"`
					HASBAGGAGE        string `xml:"HAS_BAGGAGE"`
					IBESTATUS         string `xml:"IBE_STATUS"`
					ISCASERIES        string `xml:"IS_CA_SERIES"`
					ISCHILD           string `xml:"IS_CHILD"`
					ISFIRSTTHREEROW   string `xml:"IS_FIRST_THREEROW"`
					ISGROUPMEMBER     string `xml:"IS_GROUP_MEMBER"`
					ISINIT            string `xml:"IS_INIT"`
					ISINTERLINING     string `xml:"IS_INTER_LINING"`
					INTERNATIONALFLAG string `xml:"INTERNATIONAL_FLAG"`
					ISMAIN            string `xml:"IS_MAIN"`
					ISSELECTSEAT      string `xml:"IS_SELECT_SEAT"`
					ISSUPPORTBARCODE  string `xml:"IS_SUPPORT_BARCODE"`
					ISZHCARRIER       string `xml:"IS_ZH_CARRIER"`
					NEWFARE           string `xml:"NEW_FARE"`
					ORGTERM           string `xml:"ORG_TERM"`
					PASSANGERSTATUS   string `xml:"PASSANGER_STATUS"`
					PASSENGERNAME     string `xml:"PASSENGER_NAME"`
					PNR               string `xml:"PNR"`
					SEATNO            string `xml:"SEAT_NO"`
					STATUS            string `xml:"STATUS"`
					TICKETAMOUNT      string `xml:"TICKET_AMOUNT"`
					TKTNUM            string `xml:"TKT_NUM"`
					TOCITY            string `xml:"TO_CITY"`
					TOCITYAIRPORT     string `xml:"TO_CITY_AIRPORT"`
					TOCITYNAME        string `xml:"TO_CITY_NAME"`
					TOURINDEX         string `xml:"TOUR_INDEX"`
				} `xml:"FLIGHT"`
				MESSAGE string `xml:"MESSAGE"`
			} `xml:"QUERY_TRIP_RESULT"`
		} `xml:"queryTripResponse"`
	} `xml:"Body"`
}

func TestUnmarshal(t *testing.T) {
	xmlStr := `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><ns2:queryTripResponse xmlns:ns2="http://com/shenzhenair/mobilewebservice/booking"><QUERY_TRIP_RESULT><OP_RESULT>0</OP_RESULT><FLIGHT><AIRLINE_CODE>ZH</AIRLINE_CODE><ARRIVE_DATE>2020-10-13</ARRIVE_DATE><ARRIVE_TIME>14:10</ARRIVE_TIME><BIRTH_DAY>1963-12-14</BIRTH_DAY><CANCEL_CHECK_IN_FLAG>1</CANCEL_CHECK_IN_FLAG><CARR_AIRLINE_CODE>ZH</CARR_AIRLINE_CODE><CERT_NO>4794210837794</CERT_NO><CERT_TYPE>TN</CERT_TYPE><CHECK_IN_FLAG>1</CHECK_IN_FLAG><CHECK_IN_PLAT>B2C</CHECK_IN_PLAT><CLASS_CODE>S</CLASS_CODE><CLASS_LEVEL>Y</CLASS_LEVEL><IS_AMERICAN>0</IS_AMERICAN><DST_TERM></DST_TERM><EDI_CKI>0</EDI_CKI><FLIGHT_DATE>2020-10-13</FLIGHT_DATE><FLIGHT_NO>ZH8829</FLIGHT_NO><FLIGHT_NUMBER>8829</FLIGHT_NUMBER><FLIGHT_TIME>12:20</FLIGHT_TIME><FROM_CITY>KHN</FROM_CITY><FROM_CITY_AIRPORT>南昌昌北机场</FROM_CITY_AIRPORT><IS_OPEN_CITY>0</IS_OPEN_CITY><FROM_CITY_NAME>南昌</FROM_CITY_NAME><HAS_BAGGAGE>0</HAS_BAGGAGE><IBE_STATUS>CHECKED IN</IBE_STATUS><IS_CA_SERIES>0</IS_CA_SERIES><IS_CHILD>false</IS_CHILD><IS_FIRST_THREEROW>false</IS_FIRST_THREEROW><IS_GROUP_MEMBER>false</IS_GROUP_MEMBER><IS_INIT>true</IS_INIT><IS_INTER_LINING>0</IS_INTER_LINING><INTERNATIONAL_FLAG>1</INTERNATIONAL_FLAG><IS_MAIN>false</IS_MAIN><IS_SELECT_SEAT>true</IS_SELECT_SEAT><IS_SUPPORT_BARCODE>1</IS_SUPPORT_BARCODE><IS_ZH_CARRIER>1</IS_ZH_CARRIER><NEW_FARE>1</NEW_FARE><ORG_TERM></ORG_TERM><PASSANGER_STATUS>已值机</PASSANGER_STATUS><PASSENGER_NAME>万林发</PASSENGER_NAME><PNR>PB0W1Q</PNR><SEAT_NO>25C</SEAT_NO><STATUS>CHECKED IN</STATUS><TICKET_AMOUNT>440.0</TICKET_AMOUNT><TKT_NUM>479-4210837794</TKT_NUM><TO_CITY>ZUH</TO_CITY><TO_CITY_AIRPORT>珠海金湾机场</TO_CITY_AIRPORT><TO_CITY_NAME>珠海</TO_CITY_NAME><TOUR_INDEX>1</TOUR_INDEX></FLIGHT><MESSAGE></MESSAGE></QUERY_TRIP_RESULT></ns2:queryTripResponse></soap:Body></soap:Envelope>`
	var resp queryTripResp
	if err := xml.Unmarshal([]byte(xmlStr), &resp); err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}
