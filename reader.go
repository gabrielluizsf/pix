package pix

import "errors"

// Read generates a Options struct using a copyPaste PIX code
func Read(copyPaste string) (Options, error) {
	data := buildUsingGuideMap(copyPaste, buildDataMap(Options{}))
	return readDataMap(data)
}

type intMap map[int]any

func readDataMap(data intMap) (op Options, err error) {
	keyMap, ok := data[26].(intMap)
	if !ok {
		return op, errors.New("data[26] is not (intMap)")
	}
	txMap := data[62].(intMap)
	if txMap[5].(string) == "***" {
		txMap[5] = ""
	}

	op = Options{
		Key:           keyMap[1].(string),
		Description:   keyMap[2].(string),
		Amount:        data[54].(float64),
		Name:          data[59].(string),
		City:          data[60].(string),
		TransactionID: txMap[5].(string),
	}

	return op, err
}
