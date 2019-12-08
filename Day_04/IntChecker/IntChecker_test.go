package IntChecker

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	newTests := []newTest{
		{123456, IntChecker{123456, []int{1, 2, 3, 4, 5, 6}}},
		{0, IntChecker{0, []int{0, 0, 0, 0, 0, 0}}},
	}
	for _, test := range newTests {
		testNew(t, test)
	}
}

func TestIntChecker_Check(t *testing.T) {
	checkTests := []boolTest{
		{IntChecker{111111, []int{1, 1, 1, 1, 1, 1}}, true},
		{IntChecker{223450, []int{2, 2, 3, 4, 5, 0}}, false},
		{IntChecker{123789, []int{1, 2, 3, 7, 8, 9}}, false},
		{IntChecker{112233, []int{1, 1, 2, 2, 3, 3}}, true},
		{IntChecker{123444, []int{1, 2, 3, 4, 4, 4}}, true},
		{IntChecker{111122, []int{1, 1, 1, 1, 2, 2}}, true},
	}
	for _, test := range checkTests {
		testIntCheckerCheck(t, test)
	}
}

func TestIntChecker_FullCheck(t *testing.T) {
	fullCheckTests := []boolTest{
		{IntChecker{111111, []int{1, 1, 1, 1, 1, 1}}, true},
		{IntChecker{223450, []int{2, 2, 3, 4, 5, 0}}, false},
		{IntChecker{123789, []int{1, 2, 3, 7, 8, 9}}, false},
		{IntChecker{112233, []int{1, 1, 2, 2, 3, 3}}, true},
		{IntChecker{123444, []int{1, 2, 3, 4, 4, 4}}, false},
		{IntChecker{111122, []int{1, 1, 1, 1, 2, 2}}, true},
	}
	for _, test := range fullCheckTests {
		testIntCheckerFullCheck(t, test)
	}
}

func TestCheck(t *testing.T) {
	checkTests2 := []checkTest2{
		{IntChecker{111111, []int{1, 1, 1, 1, 1, 1}}, repeats, true},
		{IntChecker{223450, []int{2, 2, 3, 4, 5, 0}}, repeats, false},
		{IntChecker{123789, []int{1, 2, 3, 7, 8, 9}}, repeats, false},
		{IntChecker{112233, []int{1, 1, 2, 2, 3, 3}}, repeats, true},
		{IntChecker{123444, []int{1, 2, 3, 4, 4, 4}}, repeats, true},
		{IntChecker{111122, []int{1, 1, 1, 1, 2, 2}}, repeats, true},
		{IntChecker{111111, []int{1, 1, 1, 1, 1, 1}}, fullRepeat, true},
		{IntChecker{223450, []int{2, 2, 3, 4, 5, 0}}, fullRepeat, false},
		{IntChecker{123789, []int{1, 2, 3, 7, 8, 9}}, fullRepeat, false},
		{IntChecker{112233, []int{1, 1, 2, 2, 3, 3}}, fullRepeat, true},
		{IntChecker{123444, []int{1, 2, 3, 4, 4, 4}}, fullRepeat, false},
		{IntChecker{111122, []int{1, 1, 1, 1, 2, 2}}, fullRepeat, true},
		{IntChecker{12345, []int{0, 1, 2, 3, 4, 5}}, repeats, false},
		{IntChecker{1234567, []int{12, 3, 4, 5, 6, 7}}, repeats, false},
	}
	for _, test := range checkTests2 {
		testCheck(t, test)
	}
}

type newTest struct {
	input    int
	expected IntChecker
}

type boolTest struct {
	input    IntChecker
	expected bool
}

type checkTest2 struct {
	input    IntChecker
	function func(IntChecker) bool
	expected bool
}

func testNew(t *testing.T, test newTest) {
	output := New(test.input)
	fmt.Printf("%v %v", test, output)
	if !reflect.DeepEqual(*output, test.expected) {
		t.Errorf("Unexpected output: expected (%d) got (%d)", test.expected, output)
	}
}

func testIntCheckerCheck(t *testing.T, test boolTest) {
	output := test.input.Check()
	if output != test.expected {
		t.Errorf("Unexpected output: expected (%t) got (%t)", test.expected, output)
	}
}

func testIntCheckerFullCheck(t *testing.T, test boolTest) {
	output := test.input.FullCheck()
	if output != test.expected {
		t.Errorf("Unexpected output: expected (%t) got (%t)", test.expected, output)
	}
}

func testCheck(t *testing.T, test checkTest2) {
	output := test.input.check(test.function)
	if output != test.expected {
		t.Errorf("Unexpected output: expected (%t) got (%t)", test.expected, output)
	}
}
