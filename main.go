package main

import (
	"fmt"
	"log"
	"math"
)

type norgateMathError struct {
	menor0 bool
	err    error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error ocurred: \n O núero é < 0:  %v \n %v", n.menor0, n.err)
}

func main() {
	var f float64 = 10

	v, err := sqrt(f)
	if err != nil {
		n := norgateMathError{false, nil}
		MenorMaior(f, n)
		log.Println(err)
	} else {
		fmt.Println("O valor da raiz de: ", f, " é ", v)
	}
}

func MenorMaior(f float64, n norgateMathError) {
	if f < 0 {
		n.menor0 = true
	} else {
		n.menor0 = false
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, norgateMathError{true, nme}
	}
	return math.Sqrt(f), nil
}
