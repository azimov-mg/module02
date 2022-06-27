package main

import "fmt"

func main() {
	Weekends, toWeekdays := make([]string, 2, 2), []string{
		"Monday", "Tuesday", "Wensday", "Thursday", "Friday", "Saturday", "Sunday",
	}
	// fmt.Println("toWeekdays -", len(toWeekdays), cap(toWeekdays))
	// fmt.Println("Weekends -", len(Weekends), cap(Weekends))

	_ = copy(Weekends, toWeekdays[5:7])
	toWeekdays = toWeekdays[:5]
	fmt.Println("Weekdays -", toWeekdays)
	fmt.Println("Weekends -", Weekends)

	allDays := append(toWeekdays, Weekends...)
	fmt.Println("All days of week -", allDays)

	// fmt.Println("toWeekdays -", len(toWeekdays), cap(toWeekdays))
	// fmt.Println("Weekends -", len(Weekends), cap(Weekends))
}
