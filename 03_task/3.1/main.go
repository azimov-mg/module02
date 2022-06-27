package main

import "fmt"

func main() {
	Weekends, toWeekdays := make([]string, 5, 5), []string{
		"Monday", "Tuesday", "Wensday", "Thursday", "Friday", "Saturday", "Sunday",
	}
	_ = copy(Weekends, toWeekdays[5:7])
	toWeekdays = toWeekdays[:5]
	fmt.Println("Weekdays -", toWeekdays)
	fmt.Println("Weekends -", Weekends)
}
