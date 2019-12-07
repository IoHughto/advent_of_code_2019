package main

import (
	"errors"
	"github.com/IoHughto/advent_of_code_2019/Day_03/wire"
	"reflect"
	"testing"
)

func TestGetWireData(t *testing.T) {
	getWireDataTests := []getWireDataTest{
		{
			input:    "R123,D234\nL345,U456",
			expected: []string{"R123,D234", "L345,U456"},
			errType:  nil,
		},
		{
			input:    "R1",
			expected: []string{},
			errType:  errOneWire,
		},
		{
			input:    "",
			expected: []string{},
			errType:  errOneWire,
		},
	}
	for _, test := range getWireDataTests {
		testGetWireData(t, test)
	}
}

func TestGetWires(t *testing.T) {
	getWiresTests := []getWiresTest{}
	for _, test := range getWiresTests {
		testGetWires(t, test)
	}
}

func TestTraverseWires(t *testing.T) {
	traverseWiresTests := []traverseWiresTest{}
	for _, test := range traverseWiresTests {
		testTraverseWires(t, test)
	}
}

func TestGetClosestIntersection(t *testing.T) {
	getClosestIntersectionTests := []getClosestIntersectionTest{}
	for _, test := range getClosestIntersectionTests {
		testGetClosestIntersection(t, test)
	}
}

type getWireDataTest struct {
	input    string
	expected []string
	errType  error
}

type getWiresTest struct {
	input    []string
	expected []wire.Wire
	errType  error
}

type traverseWiresTest struct {
	input    []wire.Wire
	expected []wire.Wire
	errType  error
}

type getClosestIntersectionTest struct {
	input      []wire.Intersection
	lengthType string
	expected   wire.Intersection
	errType    error
}

func testGetWireData(t *testing.T, test getWireDataTest) {
	output, err := getWireData(test.input)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
	}
	for i := range output {
		if output[i] != test.expected[i] {
			t.Errorf("Unexpected output: expected (%#v) got (%#v)", test.expected[i], output[i])
		}
	}
}

func testGetWires(t *testing.T, test getWiresTest) {
	output, err := getWires(test.input)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
	}
	if reflect.DeepEqual(output, test.expected) {
		t.Errorf("Unexpected output: expected (%v) got (%v)", test.expected, output)
	}
}

func testTraverseWires(t *testing.T, test traverseWiresTest) {
	err := traverseWires(test.input)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
	}
	if reflect.DeepEqual(test.input, test.expected) {
		t.Errorf("Unexpected output: expected (%v) got (%v)", test.expected, test.input)
	}
}

func testGetClosestIntersection(t *testing.T, test getClosestIntersectionTest) {
	output, err := getClosestIntersection(test.input, test.lengthType)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
	}
	if reflect.DeepEqual(output, test.expected) {
		t.Errorf("Unexpected output: expected (%v) got (%v)", test.expected, output)
	}
}
