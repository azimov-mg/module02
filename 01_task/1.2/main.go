package main

import "fmt"

func main() {
	type (
		AmericanVelocity float64
		EuropeanVelocity float64
	)

	var (
		FSpeed AmericanVelocity
		SSpeed EuropeanVelocity
	)

	const (
		fSpeed = 120.4
		sSpeed = 130
	)

	SSpeed = EuropeanVelocity(fSpeed * 3.6)
	FSpeed = AmericanVelocity(sSpeed * 2.23)

	fmt.Println(SSpeed, FSpeed)
}
