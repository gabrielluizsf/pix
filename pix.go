package pix

// New generates a Copy and Paste Pix code
func New(options Options) (string, error) {
	if err := validateData(options); err != nil {
		return "", err
	}

	data := buildDataMap(options)
	str := parseData(data)

	str += "6304"

	crc, err := calculateCRC16(str)

	if err != nil {
		return "", err
	}

	str += crc

	return str, nil
}
