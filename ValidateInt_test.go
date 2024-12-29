package main

import (
	"errors"
	"testing"
)

var tests = []struct {
	input    string
	encoding string
	expected int
	err error

}{
	{"1234", "utf8", 1234, nil},
	{"abac", "utf8", 0, errors.New("Not a number")},
	{"-1000", "utf8", -1000, nil},
	{"1-2", "utf8", 0, errors.New("Not a number")},
	{"0", "utf8", 0, nil},
	{"1234asdf", "utf8", 0, errors.New("Not a number")},
	{"asdf12341234", "utf8", 0, errors.New("Not a number")},
	{"\n\n\n\n\n\n\n\n\n", "utf8", 0, errors.New("Not a number")},
	{"jopa_slona5", "utf8", 0, errors.New("Not a number")},
	{"1234567890", "utf8", 1234567890, nil},
}

func TestValidateInt(t *testing.T) {

	for _, testcase := range tests {
		t.Run(testcase.input, func(t *testing.T) {
			got, err := ValidateInt(testcase.input, testcase.encoding)
			if err != testcase.err{
				t.Errorf("ValidateInt(%q) wrong err got %s expected %s", testcase.input, err, testcase.err)
			}
			if got != testcase.expected{
				t.Errorf("ValidateInt(%q) = %d, expected %d", testcase.input, got, testcase.expected)
			}
		})
	}
}
