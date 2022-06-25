package main

import "fmt"

func main() {
	var (
		A *int
		B int = 55
	)

	A = &B
	fmt.Println(*A)

	*A = 66
	fmt.Println(B)
}
