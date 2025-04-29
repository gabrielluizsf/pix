package pix

import (
	"bytes"
	"errors"
	"image"
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestValues_FullOptions(t *testing.T) {
	tests := []struct {
		input Options
		want  string
	}{
		{Options{
			Name:        "Gabriel Luiz",
			Key:         "7a067b11-bce7-406f-af8a-2dcf82c429d6",
			City:        "Caruaru",
			Amount:      20.67,
			Description: "Invoice #1000",
		}, "00020126750014BR.GOV.BCB.PIX01367a067b11-bce7-406f-af8a-2dcf82c429d60213Invoice #1000520400005303986540520.675802BR5912Gabriel Luiz6007Caruaru62410503***50300017BR.GOV.BCB.BRCODE01051.0.063047450"},
		{Options{
			Name:        "Gabriel Luiz",
			Key:         "1e857953-ee0f-4745-a85e-2c925dbb3c6c",
			City:        "Catende",
			Amount:      5.50,
			Description: "Lunch money",
		}, "00020126730014BR.GOV.BCB.PIX01361e857953-ee0f-4745-a85e-2c925dbb3c6c0211Lunch money52040000530398654045.505802BR5912Gabriel Luiz6007Catende62410503***50300017BR.GOV.BCB.BRCODE01051.0.063040B22"},
	}

	for _, tt := range tests {
		testValue(t, tt.input, tt.want)
	}
}

func TestValues_TransactionID(t *testing.T) {
	tests := []struct {
		input Options
		want  string
	}{
		{Options{
			Name:          "Gabriel Luiz",
			Key:           "7a067b11-bce7-406f-af8a-2dcf82c429d6",
			City:          "Caruaru",
			Amount:        20.67,
			Description:   "Invoice #5000",
			TransactionID: "341834",
		}, "00020126750014BR.GOV.BCB.PIX01367a067b11-bce7-406f-af8a-2dcf82c429d60213Invoice #5000520400005303986540520.675802BR5912Gabriel Luiz6007Caruaru6244050634183450300017BR.GOV.BCB.BRCODE01051.0.06304160B"},
	}

	for _, tt := range tests {
		testValue(t, tt.input, tt.want)
	}
}

func TestValues_WithoutAmount(t *testing.T) {
	tests := []struct {
		input Options
		want  string
	}{
		{Options{
			Name: "Gabriel Luiz",
			Key:  "1e857953-ee0f-4745-a85e-2c925dbb3c6c",
			City: "Caruaru",
		}, "00020126620014BR.GOV.BCB.PIX01361e857953-ee0f-4745-a85e-2c925dbb3c6c020052040000530398654040.005802BR5912Gabriel Luiz6007Caruaru62410503***50300017BR.GOV.BCB.BRCODE01051.0.063049D1F"},
	}

	for _, tt := range tests {
		testValue(t, tt.input, tt.want)
	}
}

func TestValues_Errors(t *testing.T) {
	tests := []struct {
		input Options
		want  error
	}{
		{Options{}, errors.New("key must not be empty")},
		{Options{
			Key:  "1e857953-ee0f-4745-a85e-2c925dbb3c6c",
			Name: "Receiver long name to cause error",
			City: "Caruaru",
		}, errors.New("name must be at least 25 characters long")},
		{Options{
			Key:  "7a067b11-bce7-406f-af8a-2dcf82c429d6",
			Name: "Gabriel Luiz",
			City: "Receiver city long name",
		}, errors.New("city must be at least 15 characters long")},
		{Options{
			Name: "Gabriel Luiz",
			Key:  "1e857953-ee0f-4745-a85e-2c925dbb3c6c",
		}, errors.New("city must not be empty")},
		{Options{
			City: "Caruaru",
			Key:  "7a067b11-bce7-406f-af8a-2dcf82c429d6",
		}, errors.New("name must not be empty")},
	}

	for _, tt := range tests {
		testError(t, tt.input, tt.want)
	}
}

// Generate a QR Code from a Pix Copy and Paste string and decode the result
func TestQrCodeContent(t *testing.T) {
	str := "00020126620014BR.GOV.BCB.PIX01361e857953-ee0f-4745-a85e-2c925dbb3c6c020052040000530398654040.005802BR5912Gabriel Luiz6007Caruaru62410503***50300017BR.GOV.BCB.BRCODE01051.0.063049D1F"
	options := QRCodeOptions{Content: str}
	qr, err := QRCode(options)
	assert.NoError(t, err)
	_, _, err = image.Decode(bytes.NewReader(qr))
	assert.NoError(t, err)
}

func testValue(t *testing.T, options Options, copyPaste string) {
	v, err := New(options)

	if err != nil {
		t.Errorf("Pix(%v) returned an error: %v", options, err)
	}
	assert.Equal(t, v, copyPaste)

	vp, err := Read(copyPaste)
	if err != nil {
		t.Errorf("ReadPix(%v) returned an error: %v", options, err)
	}
	assert.Equal(t, vp, options)
}

func testError(t *testing.T, input Options, want error) {
	_, err := New(input)
	assert.NotNil(t, err)
	assert.Equal(t, err, want)
}
