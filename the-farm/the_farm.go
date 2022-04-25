package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
	cowsNumber int
}

func (err SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", err.cowsNumber)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()

	if err == ErrScaleMalfunction && fodderAmount > 0 {
		return fodderAmount * 2 / float64(cows), nil
	}

	if err != nil && err != ErrScaleMalfunction {
		return 0.0, err
	}

	if fodderAmount < 0 {
		return 0.0, errors.New("negative fodder")
	}

	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}

	if cows < 0 {
		return 0.0, SillyNephewError{
			cowsNumber: cows,
		}
	}

	return fodderAmount / float64(cows), nil
}
