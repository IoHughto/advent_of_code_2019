package main

import (
	"errors"
	"testing"
)

func TestGetProgramData(t *testing.T) {
	dataTests := []dataTest{
		{"7,2,87", []int{7, 2, 87}, nil},
		{"7,2,87,", []int{7, 2, 87}, nil},
		{"2,notanumber,3", []int{}, errIntConversion},
	}
	for _, test := range dataTests {
		testGetProgramData(t, test)
	}
}

type dataTest struct {
	input    string
	expected []int
	errType  error
}

func testGetProgramData(t *testing.T, test dataTest) {
	data, err := getProgramData(test.input)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: %s", err)
	}
	for i, datum := range data {
		if datum != test.expected[i] {
			t.Errorf("Unexected result: expected (%v) got (%v)", test.expected, data)
		}
	}
}
