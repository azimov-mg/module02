package main

import "fmt"

func contains(a []string, x string) bool {
	for key := range a {
		if a[key] == x {
			return true
		}
	}
	return false
}

func getMax(val ...int) int {
	var max int = val[0]
	for _, v := range val {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	KeyWord, MySlice := "Banana", []string{"Apple", "Banana", "Lemon"}
	fmt.Println(contains(MySlice, KeyWord))
	fmt.Println(getMax(1, 2, 5, 2, 6, 0))
}
