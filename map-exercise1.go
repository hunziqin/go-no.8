package main

import "fmt"

func main() {
	var value int
	var isPresent bool

	day := make(map[string]int)
	day["monday"] = 1
	day["tuesday"] = 2
	day["wednesday"] = 3
	day["thursday"] = 4
	day["friday"] = 5
	day["saturday"] = 6
	day["sunday"] = 7

	for key, Value := range day {
		fmt.Printf("key is %s - Value is %d\n", key, Value)
	}

	value, isPresent = day["thursday"]
	if isPresent {
		fmt.Printf("The value of \"thursday\" in day is: %d\n", value)
	} else {
		fmt.Printf("day does not contain thursday")
	}

	value, isPresent = day["hollyday"]
	if isPresent {
		fmt.Printf("The value of \"hollyday\" in day is: %d\n", value)
	} else {
		fmt.Printf("day does not contain hollyday")
	}
}
