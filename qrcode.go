package pix

import "github.com/i9si-sistemas/qrcode"

// QRCode returns a graphical representation of the Copy and Paste code in a QR Code form.
func QRCode(options QRCodeOptions) ([]byte, error) {
	return qrcode.Encode(options.Content, qrcode.Medium, options.Size)
}
