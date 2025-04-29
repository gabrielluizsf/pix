package pix

import (
	"fmt"
	"reflect"
	"strconv"
)

func parseData(data intMap) string {
	var str string

	keys := sortKeys(data)

	for _, k := range keys {
		v := reflect.ValueOf(data[k])

		switch v.Kind() {
		case reflect.String:
			value := data[k].(string)
			str += fmt.Sprintf("%02d%02d%s", k, len(value), value)
		case reflect.Float64:
			value := strconv.FormatFloat(v.Float(), 'f', 2, 64)

			str += fmt.Sprintf("%02d%02d%s", k, len(value), value)
		case reflect.Map:
			// If the element is another map, do a recursive call
			content := parseData(data[k].(intMap))

			str += fmt.Sprintf("%02d%02d%s", k, len(content), content)
		}
	}

	return str
}
