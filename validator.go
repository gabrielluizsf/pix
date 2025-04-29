package pix

import (
	"errors"
	"unicode/utf8"
)

func validateData(options Options) error {
	if options.Key == "" {
		return errors.New("key must not be empty")
	}

	if options.Name == "" {
		return errors.New("name must not be empty")
	}

	if options.City == "" {
		return errors.New("city must not be empty")
	}

	if utf8.RuneCountInString(options.Name) > 25 {
		return errors.New("name must be at least 25 characters long")
	}

	if utf8.RuneCountInString(options.City) > 15 {
		return errors.New("city must be at least 15 characters long")
	}

	return nil
}
