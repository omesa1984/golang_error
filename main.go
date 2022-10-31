package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	var f float64 = -10
	v, err := sqrt(f)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("O valor da raiz de: ", f, " é ", v)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	return math.Sqrt(f), nil
}
