/*
	* ----- ERROR

	! zero value : nil
*/

package main

import (
	"errors"
	"fmt"
)

// OLD function
// func Devide(a, b float64) float64 {
// 	r := a / b
// 	return r
// }

type myError struct{}

func (e myError) Error() string {
	return "My Error"
}

// 2.declare error
var myErr = errors.New("HERE IS ERROR")

// Function with handle Error
func Devide(a, b float64) (float64, error) {
	if b == 0 {
		// err := myError{}

		// 1.declare error
		// err := fmt.Errorf("Can't devide '%.2f' by zero", a)
		// return 0, err

		// 2.
		return 0, myErr

	}
	r := a / b
	return r, nil
}

func main() {
	r, e := Devide(1, 0)
	//  also declare e for error
	if e != nil {
		// logic
		fmt.Println("handle this Error: ", e)
		return
	}
	fmt.Println(r, e)

}
