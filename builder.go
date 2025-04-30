package pix

import (
	"reflect"
	"strconv"
)

func buildDataMap(options Options) intMap {
	data := make(intMap)

	// Payload Format Indicator
	data[0] = "01"

	// Merchant Account Information
	data[26] = intMap{0: "BR.GOV.BCB.PIX", 1: options.Key, 2: options.Description}

	// Merchant Category Code
	data[52] = "0000"

	// Transaction Currency - Brazilian Real - ISO4217
	data[53] = "986"

	// Transaction Amount
	data[54] = options.Amount

	// Country Code - ISO3166-1 alpha 2
	data[58] = "BR"

	// Merchant Name. 25 characters maximum
	data[59] = options.Name

	// Merchant City. 15 characters maximum
	data[60] = options.City

	// Transaction ID
	data[62] = intMap{5: "***", 50: intMap{0: "BR.GOV.BCB.BRCODE", 1: "1.0.0"}}

	if options.TransactionID != "" {
		data[62].(intMap)[5] = options.TransactionID
	}

	return data
}

func buildUsingGuideMap(copyPaste string, guide intMap) intMap {
	data := make(intMap)

	k := 0
	convert := func(cp string, k *int, iv int) int {
		value, _ := strconv.Atoi(cp[*k : *k+iv])
		*k += 2
		return value
	}
	for k < len(copyPaste) {
		index := convert(copyPaste, &k, 2)
		length := convert(copyPaste, &k, 2)

		value := copyPaste[k : k+length]
		k += length

		v := reflect.ValueOf(guide[index])
		switch v.Kind() {
		case reflect.Map:
			m := guide[index].(intMap)
			data[index] = buildUsingGuideMap(value, m)
		case reflect.String:
			data[index] = value
		case reflect.Float64:
			data[index], _ = strconv.ParseFloat(value, 64)
		}
	}

	return data
}
