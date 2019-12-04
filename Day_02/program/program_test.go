package program

import (
	"errors"
	"testing"
)

func TestIntCode_Run(t *testing.T) {
	runTests := []runTest{
		{[]int{1, 0, 0, 0, 99}, 2, []int{2, 0, 0, 0, 99}, nil},
		{[]int{2, 3, 0, 3, 99}, 2, []int{2, 3, 0, 6, 99}, nil},
		{[]int{2, 4, 4, 5, 99, 0}, 2, []int{2, 4, 4, 5, 99, 9801}, nil},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, 30, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, nil},
		{[]int{4, 0, 0, 0}, 0, []int{}, errInvalidInstruction},
		{[]int{1, 0, 0, 0}, 0, []int{}, errEndOfProgram},
	}
	for _, test := range runTests {
		testIntCodeRun(t, test)
	}
}

func TestIntCode_Replace(t *testing.T) {
	replaceTests := []replaceTest{
		{
			[]int{0, 1, 2, 3, 4},
			[]ReplaceData{{0, 10}, {1, 11}},
			[]int{10, 11, 2, 3, 4},
			nil,
		},
		{
			[]int{0, 1, 2, 3, 4},
			[]ReplaceData{{5, 10}},
			[]int{},
			errInvalidPosition,
		},
		{
			[]int{0, 1, 2, 3, 4},
			[]ReplaceData{{500, 10}},
			[]int{},
			errInvalidPosition,
		},
	}
	for _, test := range replaceTests {
		testIntCodeReplace(t, test)
	}
}

type runTest struct {
	input           []int
	expectedResult  int
	expectedProgram []int
	errType         error
}

type replaceTest struct {
	input    []int
	replaces []ReplaceData
	expected []int
	errType  error
}

func testIntCodeRun(t *testing.T, test runTest) {
	program := IntCode{}
	program.Load(test.input)
	result, err := program.Run()
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: %s", err)
	}
	if err != nil {
		return
	}
	if result != test.expectedResult {
		t.Errorf("Unexected result: expected (%d) got (%d)", test.expectedResult, result)
	}
	for i := range program.data {
		if program.data[i] != test.expectedProgram[i] {
			t.Errorf("Unexected result: expected (%v) got (%v)", test.expectedProgram, program.data)
		}
	}
}

func testIntCodeReplace(t *testing.T, test replaceTest) {
	program := IntCode{}
	program.Load(test.input)
	program, err := program.Replace(test.replaces)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: %s", err)
	}
	if err != nil {
		return
	}
}
