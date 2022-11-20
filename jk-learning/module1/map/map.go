package main

import "fmt"

func main() {
	m3 := map[string]int{"王强": 39, "小李": 45, "芳芳": 35}
	for k, v := range m3 {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
