package wire

import (
	"fmt"
	"strconv"
)

type Wire struct {
	Steps []step
	Path  map[Position]bool
}

type step struct {
	direction rune
	length    int
}

type Position struct {
	X int
	Y int
}

type Intersection struct {
	Position Position
	Length   int
}

func (w *Wire) Load(stepData []string) error {
	for _, stepDatum := range stepData {
		step, err := parseStep(stepDatum)
		if err != nil {
			return err
		}
		w.Steps = append(w.Steps, step)
	}
	w.Path = make(map[Position]bool)
	return nil
}

func (w *Wire) Traverse() error {
	currentPosition := Position{0, 0}
	w.Path[currentPosition] = true
	for _, step := range w.Steps {
		for i := 0; i < step.length; i++ {
			err := currentPosition.move(step.direction)
			if err != nil {
				return err
			}
			w.Path[currentPosition] = true
		}
	}
	return nil
}

func (p *Position) move(direction rune) error {
	switch direction {
	case 'R':
		p.right()
	case 'L':
		p.left()
	case 'U':
		p.up()
	case 'D':
		p.down()
	default:
		return fmt.Errorf("%w: %s", errInvalidDirection, string(direction))
	}
	return nil
}

func (w *Wire) Intersect(ow Wire) ([]Intersection, error) {
	var intersections []Intersection
	for testPosition := range w.Path {
		if ow.Path[testPosition] {
			intersections = append(intersections, Intersection{testPosition, 0})
		}
	}

	for i, intersection := range intersections {
		wLength, err := w.DistanceAlong(intersection.Position)
		if err != nil {
			return nil, err
		}

		owLength, err := ow.DistanceAlong(intersection.Position)
		if err != nil {
			return nil, err
		}
		intersections[i].Length = wLength + owLength
	}

	return intersections, nil
}

func (w *Wire) DistanceAlong(position Position) (int, error) {
	length := 0
	currentPosition := Position{}
	for _, step := range w.Steps {
		for i := 0; i < step.length; i++ {
			length++
			err := currentPosition.move(step.direction)
			if err != nil {
				return 0, err
			}
			if currentPosition == position {
				return length, nil
			}
		}
	}
	if currentPosition == position {
		return length, nil
	}
	return 0, errNotAlongPath
}

func (i *Intersection) GetLength(lengthType string) (int, error) {
	switch lengthType {
	case "manhattan":
		return i.Position.manhattanDistance(), nil
	case "along":
		return i.Length, nil
	default:
		return 0, fmt.Errorf("%w: %s", errInvalidLengthType, lengthType)
	}
}

func parseStep(s string) (step, error) {
	direction := []rune(s[0:1])[0] // this is gross
	length, err := strconv.Atoi(s[1:])
	if err != nil {
		return step{}, fmt.Errorf("%w, %s", errIntParse, err)
	}
	return step{direction, length}, nil
}

func (p *Position) right() {
	p.X++
}

func (p *Position) left() {
	p.X--
}

func (p *Position) up() {
	p.Y++
}

func (p *Position) down() {
	p.Y--
}

func (p *Position) manhattanDistance() int {
	return abs(p.X) + abs(p.Y)
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

// sentinel errors
var errInvalidDirection = fmt.Errorf("invalid direction")
var errNotAlongPath = fmt.Errorf("intersection not along path of wire")
var errInvalidLengthType = fmt.Errorf("invalid length type")
var errIntParse = fmt.Errorf("couldn't parse int")
