package pix

// New generates a Copy and Paste Pix code
func New(options Options) (copyPaste string,  err error) {
	if err := validateData(options); err != nil {
		return "", err
	}

	data := buildDataMap(options)
	copyPaste = parseData(data)

	copyPaste += "6304"

	crc, err := calculateCRC16(copyPaste)

	if err != nil {
		return "", err
	}

	copyPaste += crc

	return copyPaste, nil
}
