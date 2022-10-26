package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	var extractedValue int
	switch fnb.(type) {
	case FancyNumber:
		extractedValue, _ = strconv.Atoi(fnb.Value())
	default:
		extractedValue = 0
	}
	return extractedValue
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	numba := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(numba))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	var returnString string
	switch v := i.(type) {
	case int:
		returnString = DescribeNumber(float64(v))
	case float64:
		returnString = DescribeNumber(v)
	case NumberBox:
		returnString = DescribeNumberBox(v)
	case FancyNumberBox:
		returnString = DescribeFancyNumberBox(v)
	default:
		returnString = "Return to sender"
	}
	return returnString
}
