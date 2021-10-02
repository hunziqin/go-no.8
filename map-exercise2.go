package main

import (
	"fmt"
	"sort"
)

var (
	drink = map[string]string{"cocacola": "可口可乐", "sprite": "雪碧",
		"yuanqi forest": "元气森林", "pepsi": "百事可乐"}
)

func main() {
	for k, _ := range drink {
		fmt.Printf("%s\n", k)
	}
	for k, v := range drink {
		fmt.Printf("%s %s\n", k, v)
	}
	keys := make([]string, len(drink))
	i := 0
	for k, _ := range drink {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s %s / ", k, drink[k])
	}
}
