package errors

import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f2(number int) (int, error) {
	if number == 42 {
		return -1, &argError{number, "Can not work with it"}
	}
	return number + 3, nil
}

func PrintCustomError() {
	_, err := f2(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}

	_, err2 := f2(44)
	var ae2 *argError
	if errors.As(err2, &ae2) {
		fmt.Println(ae2.arg)
		fmt.Println(ae2.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
