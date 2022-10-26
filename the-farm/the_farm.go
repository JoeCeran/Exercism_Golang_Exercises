package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNewphewError struct {
	cows int
}

func (e *SillyNewphewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fod, err := weightFodder.FodderAmount()
	var returnedFod float64

	if cows == 0 {
		err = errors.New("division by zero")
	} else if cows < 0 {
		err = &SillyNewphewError{cows: cows}
	}

	if fod > 0 && err == ErrScaleMalfunction || fod > 0 && err == nil {
		if err == ErrScaleMalfunction {
			fod = fod * 2
			err = nil
		}
		returnedFod = fod / float64(cows)
	} else if fod < 0 {
		if err == ErrScaleMalfunction || err == nil {
			err = errors.New("negative fodder")
		}
		returnedFod = 0
	}
	return returnedFod, err
}
