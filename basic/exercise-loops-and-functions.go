package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(Sqrt(2))
	}
}
