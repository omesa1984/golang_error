package main

import (
	"fmt"
	"log"
	"math"
)

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
		return 0, fmt.Errorf("norgate math: square root of negative number: %v", f)
	}
	return math.Sqrt(f), nil
}
