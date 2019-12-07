package program

import "fmt"

type IntCode struct {
	data []int
}

type ReplaceData struct {
	Position int
	Value    int
}

const (
	ADD_CODE      = 1
	MULTIPLY_CODE = 2
	END_CODE      = 99
)

func (ic *IntCode) Load(data []int) {
	ic.data = append([]int{}, data...)
}

func (ic *IntCode) Replace(replaceData []ReplaceData) (IntCode, error) {
	for _, replaceDatum := range replaceData {
		if len(ic.data) <= replaceDatum.Position {
			return IntCode{}, fmt.Errorf("%w: %d", errInvalidPosition, replaceDatum.Position)
		}
		ic.data[replaceDatum.Position] = replaceDatum.Value
	}

	return *ic, nil
}

func (ic *IntCode) Run() (int, error) {
	position := 0
	for {
		done, err := ic.iterate(position)
		if err != nil {
			return 0, err
		}
		if done {
			break
		}
		if len(ic.data) <= position+4 {
			return 0, errEndOfProgram
		}
		position += 4
	}
	return ic.data[0], nil
}

func (ic *IntCode) iterate(position int) (bool, error) {
	switch ic.data[position] {
	case ADD_CODE:
		ic.data[ic.data[position+3]] = ic.data[ic.data[position+1]] + ic.data[ic.data[position+2]]
	case MULTIPLY_CODE:
		ic.data[ic.data[position+3]] = ic.data[ic.data[position+1]] * ic.data[ic.data[position+2]]
	case END_CODE:
		return true, nil
	default:
		return false, fmt.Errorf("%w: %d -> %d", errInvalidInstruction, position, ic.data[position])
	}
	return false, nil
}

// sentinel errors
var errInvalidInstruction = fmt.Errorf("invalid instruction type")
var errEndOfProgram = fmt.Errorf("reached end of program unexpectedly")
var errInvalidPosition = fmt.Errorf("position out of range")
