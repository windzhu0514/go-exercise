package structassgin

import (
	"bytes"
	"encoding/json"
	"reflect"
	"sort"

	"github.com/ChimeraCoder/gojson"
)

var buf bytes.Buffer

func Generate(jsonStr, structName string) string {
	buf.WriteString("var data " + structName)

	var data interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return ""
	}

	switch v := data.(type) {
	case map[string]interface{}:
		walk(v, "data")
	case []interface{}:
		buf.WriteString("data=nil")
	}

	return buf.String()
}

func walk(obj map[string]interface{}, parent string) {
	keys := make([]string, 0, len(obj))
	for key := range obj {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, key := range keys {
		value := obj[key]
		switch v := value.(type) {
		case map[string]interface{}:
			walk(v, parent+"."+gojson.FmtFieldName(key))
			continue
		case []interface{}:

		}

		buf.WriteString("\n" + parent + "." + gojson.FmtFieldName(key) + "=" + zeroValue(value))
	}
}
func zeroValue(v interface{}) string {
	tv := reflect.TypeOf(v)
	switch tv.Kind() {
	case reflect.Bool:
		return "false"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:
		return "0"
	case reflect.Map:
		return "nil"
	case reflect.Slice:
		return "nil"
	case reflect.String:
		return "\"\""
	case reflect.Struct:
		return "nil"
	}

	return ""
}
