package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
)

func main() {
	htmls := []string{``, ``, ``}
	mapFlightInfo := make(map[string]string)
	mapFlightInfo["psgrName"] = "111111111111"
	mapFlightInfo["certType"] = "22222"
	var fi flightInfo

	//if err := mapstructure.Decode(mapFlightInfo, &fi); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(fi)

	config := &mapstructure.DecoderConfig{
		DecodeHook: customHook(),
		Result:     &fi,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := decoder.Decode(mapFlightInfo); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fi)

}

func customHook() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Kind,
		t reflect.Kind,
		data interface{}) (interface{}, error) {
		_ = f
		_ = t
		_ = data

		if f == reflect.Map {
			dataVal := reflect.ValueOf(data)
			newDataVal := make(map[interface{}]interface{})
			for _, rawMapKey := range dataVal.MapKeys() {
				rawMapVal := dataVal.MapIndex(rawMapKey)
				newMapKeyString := rawMapKey.String()
				if index := strings.Index(newMapKeyString, "_"); index > 0 {
					if index2 := strings.Index(newMapKeyString[index+1:], "_"); index2 > 0 {
						newMapKeyString = newMapKeyString[index+1 : index+index2]
					} else {
						newMapKeyString = newMapKeyString[index+1:]
					}
				}

				newMapKeyString = strings.TrimSuffix(newMapKeyString, "11")
				newDataVal[newMapKeyString] = rawMapVal.Interface()
			}

			return newDataVal, nil
		}

		return data, nil
	}
}

type flightInfo struct {
	PsgrName string `mapstructure:"psgrName"`
	CertType string `mapstructure:"certType"`
	CertNo   string `mapstructure:"cert_no"`
	OrgCity  string `mapstructure:"org_city"`
}
