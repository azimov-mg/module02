package main

import (
	"fmt"
	"math"
)

func main() {
	L, R := 35, new(float64)
	*R = float64(L) / (2 * math.Phi)
	fmt.Println("Radius of Circle = ", math.Round(*R*100)/100)

	var S = math.Phi * math.Pow(*R, 2)
	fmt.Println("Square of Circle = ", math.Round(S*100)/100)
}
