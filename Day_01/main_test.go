package main

import (
	"errors"
	"testing"
)

func TestParseData(t *testing.T) {
	parseTests := []parseTest{
		{[]byte("1\n2\n3\n"), []int64{1, 2, 3}, nil},
		{[]byte("1\n2\nasdf"), []int64{1, 2, 3}, errInvalidSyntax},
	}
	for _, test := range parseTests {
		testParseData(t, test)
	}
}

func TestGetSumOfData(t *testing.T) {
	sumTests := []sumTest{
		{[]int64{1, 2, 3, 4}, 0, nil},
		{[]int64{96, 18, 21, 22}, 44, nil},
	}
	for _, test := range sumTests {
		testGetSumOfData(t, test)
	}
}

func TestGetCompleteSumOfData(t *testing.T) {
	sumTests := []sumTest{
		{[]int64{1, 2, 3, 4}, 0, nil},
		{[]int64{96, 18, 21, 22}, 52, nil},
	}
	for _, test := range sumTests {
		testGetCompleteSumOfData(t, test)
	}
}

type parseTest struct {
	input    []byte
	expected []int64
	errType  error
}

type sumTest struct {
	input    []int64
	expected int64
	errType  error
}

func testParseData(t *testing.T, test parseTest) {
	data, err := parseData(test.input)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
	}
	for i := range data {
		if data[i] != test.expected[i] {
			t.Errorf("Unexpected Data: expected (%v) got (%v)", test.expected, data)
		}
	}
}

func testGetSumOfData(t *testing.T, test sumTest) {
	totalFuel, err := getSumOfData(test.input, getSimpleTotal)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
	}
	if totalFuel != test.expected {
		t.Errorf("Unexpected result: expected (%d) got (%d)", test.expected, totalFuel)
	}
}

func testGetCompleteSumOfData(t *testing.T, test sumTest) {
	totalFuel, err := getSumOfData(test.input, getCompleteTotal)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
	}
	if totalFuel != test.expected {
		t.Errorf("Unexpected result: expected (%d) got (%d)", test.expected, totalFuel)
	}
}
