package main

import (
	"fmt"
)

func main() {
	names := []string{"小嘛", "大枪", "撒娇"}
	fr := []float64{24.2, 31, 45}

	names = append(names, "haha")
	fr = append(fr, 67)

	for i, val := range names {
		if val == "大枪" {
			fmt.Printf("%s的体脂率是 %f\n", val, fr[i])
		}
	}
}
