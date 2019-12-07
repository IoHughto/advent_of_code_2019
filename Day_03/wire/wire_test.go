package wire

import (
	"errors"
	"reflect"
	"testing"
)

func TestAbs(t *testing.T) {
	absTests := []absTest{
		{1, 1},
		{-5, 5},
		{0, 0},
	}
	for _, test := range absTests {
		testAbs(t, test)
	}
}

func TestMove(t *testing.T) {
	moveTests := []moveTest{
		{'R', Position{1, 0}, nil},
		{'L', Position{-1, 0}, nil},
		{'U', Position{0, 1}, nil},
		{'D', Position{0, -1}, nil},
		{'X', Position{}, errInvalidDirection},
	}
	for _, test := range moveTests {
		testMove(t, test)
	}
}

func TestManhattan(t *testing.T) {
	manhattanTests := []manhattanTest{
		{Position{1, 2}, 3},
		{Position{1, -1}, 2},
		{Position{0, 0}, 0},
		{Position{-24, 36}, 60},
	}
	for _, manhattanTest := range manhattanTests {
		testManhattan(t, manhattanTest)
	}
}

func TestParse(t *testing.T) {
	parseTests := []parseTest{
		{"R123", step{'R', 123}, nil},
		{"D321", step{'D', 321}, nil},
		{"Y86415", step{'Y', 86415}, nil},
		{"ABCD", step{}, errIntParse},
	}
	for _, test := range parseTests {
		testParse(t, test)
	}
}

func TestIntersection_GetLength(t *testing.T) {
	getLengthTests := []getLengthTest{
		{Intersection{Position{1, 2}, 5}, "manhattan", 3, nil},
		{Intersection{Position{4, 3}, 234}, "along", 234, nil},
		{Intersection{Position{0, 0}, 0}, "something else", 0, errInvalidLengthType},
	}
	for _, test := range getLengthTests {
		testGetLength(t, test)
	}
}

func TestWire_DistanceAlong(t *testing.T) {
	distanceAlongTests := []distanceAlongTest{
		{Wire{[]step{{'R', 10}}, nil}, Position{8, 0}, 8, nil},
		{Wire{[]step{{'R', 8}, {'U', 5}}, nil}, Position{8, 5}, 13, nil},
		{Wire{[]step{{'L', 10}}, nil}, Position{2, 2}, 0, errNotAlongPath},
		{Wire{nil, nil}, Position{0, 0}, 0, nil},
		{Wire{[]step{{'X', 5}}, nil}, Position{}, 0, errInvalidDirection},
	}
	for _, test := range distanceAlongTests {
		testDistanceAlong(t, test)
	}
}

func TestWire_Intersect(t *testing.T) {
	intersectTests := []intersectTest{
		{
			data: Wire{
				Steps: []step{{'R', 2}, {'U', 5}},
				Path:  map[Position]bool{Position{2, 5}: true},
			},
			input: Wire{
				Steps: []step{{'U', 5}, {'R', 2}},
				Path:  map[Position]bool{Position{2, 5}: true},
			},
			expected: []Intersection{
				{
					Position: Position{2, 5},
					Length:   7,
				},
			},
			errType: nil,
		},
	}
	for _, test := range intersectTests {
		testIntersect(t, test)
	}
}

func TestWire_Traverse(t *testing.T) {
	traverseTests := []traverseTest{
		{
			input: Wire{
				Steps: []step{{'U', 2}},
				Path:  nil,
			},
			expected: Wire{
				Steps: []step{{'U', 2}},
				Path:  map[Position]bool{Position{0, 0}: true, Position{0, 1}: true, Position{0, 2}: true},
			},
			errType: nil,
		},
	}
	for _, test := range traverseTests {
		testTraverse(t, test)
	}
}

func TestWire_Load(t *testing.T) {
	loadTests := []loadTest{
		{
			input: []string{"R2", "D2"},
			expected: Wire{
				Steps: []step{{'R', 2}, {'D', 2}},
				Path:  nil,
			},
			errType: nil,
		},
	}
	for _, test := range loadTests {
		testLoad(t, test)
	}
}

type absTest struct {
	input    int
	expected int
}

type moveTest struct {
	input    rune
	expected Position
	errType  error
}

type manhattanTest struct {
	input    Position
	expected int
}

type parseTest struct {
	input    string
	expected step
	errType  error
}

type getLengthTest struct {
	data     Intersection
	input    string
	expected int
	errType  error
}

type distanceAlongTest struct {
	data     Wire
	input    Position
	expected int
	errType  error
}

type intersectTest struct {
	data     Wire
	input    Wire
	expected []Intersection
	errType  error
}

type traverseTest struct {
	input    Wire
	expected Wire
	errType  error
}

type loadTest struct {
	input    []string
	expected Wire
	errType  error
}

func testAbs(t *testing.T, test absTest) {
	value := abs(test.input)
	if value != test.expected {
		t.Errorf("Unexpected value: expected (%d) got (%d)", test.expected, value)
	}
}

func testMove(t *testing.T, test moveTest) {
	position := Position{0, 0}
	err := position.move(test.input)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
	}
	if position != test.expected {
		t.Errorf("Unexpected position: expected (%v) got (%v)", test.expected, position)
	}
}

func testManhattan(t *testing.T, test manhattanTest) {
	distance := test.input.manhattanDistance()
	if test.expected != distance {
		t.Errorf("Unexpected distance: expected (%d) got (%d)", test.expected, distance)
	}
}

func testParse(t *testing.T, test parseTest) {
	step, err := parseStep(test.input)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
	}
	if step != test.expected {
		t.Errorf("Unexpected step: expected (%v) got (%v)", test.expected, step)
	}
}

func testGetLength(t *testing.T, test getLengthTest) {
	length, err := test.data.GetLength(test.input)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
	}
	if length != test.expected {
		t.Errorf("Unexpected length: expected (%d) got (%d)", test.expected, length)
	}
}

func testDistanceAlong(t *testing.T, test distanceAlongTest) {
	distance, err := test.data.DistanceAlong(test.input)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
	}
	if distance != test.expected {
		t.Errorf("Unexpected length: expected (%d) got (%d)", test.expected, distance)
	}
}

func testIntersect(t *testing.T, test intersectTest) {
	intersections, err := test.data.Intersect(test.input)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
	}
	if reflect.DeepEqual(intersections, test.expected) {
		t.Errorf("Unexpected intersections: expected (%v) got (%v)", test.expected, intersections)
	}
}

func testTraverse(t *testing.T, test traverseTest) {
	test.input.Path = make(map[Position]bool)
	err := test.input.Traverse()
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
	}
	if !reflect.DeepEqual(test.expected, test.input) {
		t.Errorf("Unexpected paths: expected (%v) got (%v)", test.expected, test.input)
	}
}

func testLoad(t *testing.T, test loadTest) {
	test.expected.Path = make(map[Position]bool)
	wire := Wire{}
	err := wire.Load(test.input)
	if !errors.Is(err, test.errType) {
		t.Errorf("Unexpected error: expected (%s) got (%s)", test.errType, err)
	}
	if !reflect.DeepEqual(wire, test.expected) {
		t.Errorf("Unexpected wire: expected (%v) got (%v)", test.expected, wire)
	}
}
