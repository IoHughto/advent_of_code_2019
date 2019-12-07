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

func TestFindTargetInput(t *testing.T) {
	findTests := []findTest{
		{[]int{1, 0, 0, 0, 99, 15, 20}, 35, 5, 6, nil},
		{[]int{99}, 1, 0, 0, errTargetNotFound},
	}
	for _, test := range findTests {
		testFindTargetInput(t, test)
	}
}

type dataTest struct {
	input    string
	expected []int
	errType  error
}

type findTest struct {
	inputData    []int
	target       int
	expectedNoun int
	expectedVerb int
	errType      error
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

func testFindTargetInput(t *testing.T, test findTest) {
	for i := len(test.inputData); i <= 100; i++ {
		test.inputData = append(test.inputData, 99)
	}
	noun, verb, err := findTargetInput(test.inputData, test.target)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: %s", err)
	}
	if noun != test.expectedNoun {
		t.Errorf("Unexected noun: expected (%v) got (%v)", test.expectedNoun, noun)
	}
	if verb != test.expectedVerb {
		t.Errorf("Unexected noun: expected (%v) got (%v)", test.expectedVerb, verb)
	}
}
