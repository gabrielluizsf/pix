package pix

import (
	"fmt"

	crc16 "github.com/i9si-sistemas/crc-16"
)

func calculateCRC16(str string) (string, error) {
	table := crc16.MakeTable(crc16.CCITT_FALSE)

	h := crc16.New(table)
	_, err := h.Write([]byte(str))

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%04X", h.Sum16()), nil
}
