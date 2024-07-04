package main

import (
	"fmt"
)

type valError struct {
	val     int
	message string
}

func (e *valError) Error() string {
	return fmt.Sprintf(e.message, e.val)
}

func IsEven(n int) (bool, error) {
	if n%2 == 0 {
		return true, nil
	} else {
		return false, &valError{n, "the number %d is not even"}
	}
}
