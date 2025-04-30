package pix

import (
	"errors"
	"unicode/utf8"
)

func validateData(options Options) error {
	if isEmpty(options.Key) {
		return errors.New("key must not be empty")
	}

	if isEmpty(options.Name) {
		return errors.New("name must not be empty")
	}

	if isEmpty(options.City) {
		return errors.New("city must not be empty")
	}

	if !validateRuneCount(options.Name, 25) {
		return errors.New("name must be at least 25 characters long")
	}

	if !validateRuneCount(options.City, 15) {
		return errors.New("city must be at least 15 characters long")
	}

	return nil
}

func isEmpty(str string) bool {
	return len(str) == 0
}

func validateRuneCount(
	str string,
	length int,
) bool {
	return utf8.RuneCountInString(str) < length
}
