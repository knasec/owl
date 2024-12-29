package main

import (
	"errors"
	"strconv"
	"strings"
)

// ErrNotANumber is a error when input data is not a number
var ErrNotANumber = errors.New("Not a number")

// ValidateInt is a function to validate incoming data and return correct int value or error
func ValidateInt(str string, encoding string) (int, error) {
	number, err := strconv.Atoi(str)

	if encoding == "ASCII" {
		asciiCodes := strings.Split(str, " ")

		var result string
		for _, code := range asciiCodes {
			charCode, err := strconv.Atoi(code)

			if err != nil {
				return 0, ErrNotANumber
			}
			result += string(rune(charCode))
		}
	}

	if err != nil {
		return 0, ErrNotANumber
	}

	return number, nil
}
