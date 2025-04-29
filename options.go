package pix

// Options is a configuration struct.
type Options struct {
	// Pix Key (CPF/CNPJ, Email, Cellphone or Random Key)
	Key string
	// Receiver name
	Name string
	// Receiver city
	City string
	// Transaction amount
	Amount float64
	// Transaction description
	Description string
	// Transaction ID
	TransactionID string
}

// QRCodeOptions is a configuration struct.
type QRCodeOptions struct {
	// QR Code content
	Content string
	// Default: 256
	Size int
}