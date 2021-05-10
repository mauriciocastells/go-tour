/*Copy your Sqrt function from the earlier exercise and modify it to return an error value.

Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

Create a new type

type ErrNegativeSqrt float64
and make it an error by giving it a

func (e ErrNegativeSqrt) Error() string
method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".

Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop.
You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?

Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.*/

package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {

	if x > 0 {
		z := 1.0
		for i := 1; float64(i) <= 10; i++ {
			z -= (z*z - x) / (2 * z)
			//fmt.Printf("%v: %v\n",i,z)
		}
		return z, nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}

}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
