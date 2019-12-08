package IntChecker

import (
	"math"
)

type IntChecker struct {
	value  int
	digits []int
}

func New(value int) *IntChecker {
	digits := getDigits(value)
	return &IntChecker{
		value:  value,
		digits: digits,
	}
}

func (ic *IntChecker) Check() bool {
	return ic.check(repeats)
}

func (ic *IntChecker) FullCheck() bool {
	return ic.check(fullRepeat)
}

func (ic *IntChecker) check(repeatCheck func(IntChecker) bool) bool {
	if ic.value > 999999 || ic.value < 100000 {
		return false
	}

	if !ic.isMonotonic() {
		return false
	}

	if !repeatCheck(*ic) {
		return false
	}

	return true
}

func getDigits(value int) []int {
	var digits []int
	for i := 5; i >= 0; i-- {
		factor := int(math.Pow(10, float64(i)))
		digit := math.Floor(float64(value) / (float64(factor)))
		digits = append(digits, int(digit))
		value -= int(digit) * factor
	}
	return digits
}

func (ic *IntChecker) isMonotonic() bool {
	for i := 0; i < len(ic.digits)-1; i++ {
		if ic.digits[i] > ic.digits[i+1] {
			return false
		}
	}
	return true
}

func repeats(ic IntChecker) bool {
	for i := 0; i < len(ic.digits)-1; i++ {
		if ic.digits[i] == ic.digits[i+1] {
			return true
		}
	}
	return false
}

func fullRepeat(ic IntChecker) bool {
	if !repeats(ic) {
		return false
	}
	var repeatCounter []int
	for i, digit := range ic.digits {
		if i == 0 {
			repeatCounter = append(repeatCounter, 1)
		} else {
			if digit == ic.digits[i-1] {
				repeatCounter[len(repeatCounter)-1]++
			} else {
				repeatCounter = append(repeatCounter, 1)
			}
		}
	}
	for _, counter := range repeatCounter {
		if counter > 2 && math.Mod(float64(counter), 2) == 1 {
			return false
		}
	}
	return true
}
