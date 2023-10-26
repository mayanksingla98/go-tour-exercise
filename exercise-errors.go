package main

import (
	"fmt"
	"strconv"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return 0, nil
}

func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number: " + strconv.FormatFloat(float64(e), 'f', -1, 64)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-6.4543))
}
