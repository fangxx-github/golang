//输入4个点的坐标：1.定义8个flout64；2.从命令行录入
//两个点计算斜率：

package main

import "fmt"

func main() {
	var p1x, p1y, p2x, p2y, p3x, p3y, p4x, p4y float64
	fmt.Print("录入第一个点的x:")
	fmt.Scanln(&p1x)
	fmt.Print("录入第一个点的y:")
	fmt.Scanln(&p1y)

	fmt.Print("录入第二个点的x:")
	fmt.Scanln(&p2x)
	fmt.Print("录入第二个点的y:")
	fmt.Scanln(&p2y)

	fmt.Print("录入第三个点的x:")
	fmt.Scanln(&p3x)
	fmt.Print("录入第三个点的y:")
	fmt.Scanln(&p3y)

	fmt.Print("录入第四个点的x:")
	fmt.Scanln(&p4x)
	fmt.Print("录入第四个点的y:")
	fmt.Scanln(&p4y)

	k1 := (p2y - p1y) / (p2x - p1x)
	k2 := (p4y - p3y) / (p4x - p3x)

	if k1 == k2 {
		fmt.Println("两条线平行")
	} else {
		fmt.Println("两条线不平行")
	}
}
