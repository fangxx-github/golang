/*package main

import "fmt"

func main() {
	a := [3]string{"小强", "男", "在职"}
	b := [3]string{"小强", "男", "在职"}
	c := [3]string{"小强", "男", "在职"}

	newgroupinfo := [3][3]string{a, b, c}

	fmt.Println(newgroupinfo)

}

/*func main() {
	a := [3]string{"小名", "男", "在职"}
	b := [3]string{"小强", "男", "在职"}
	c := [3]string{"小东", "女", "在职"}

	fmt.Println("第一种方式普通")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("第二种以多维数组2方式")
	d := [3][3]string{a, b, c}
	for d1, d1val := range d {
		for d2, d2val := range d1val {
			fmt.Println(d1, d2, d2val)
		}
	}

}*/

func main() {
	a := [3]int{1, 2, 3}
	var b [3]int

	for i := len(a); i > 0; i-- {
		b[i-1] = a[i-1]

	}
	fmt.Println(b)
}
*/