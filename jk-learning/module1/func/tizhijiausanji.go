package main

import "fmt"

func main() {
	for {
		mainFatRateBody()

		/*cont := whetherContinue()
		if !cont {
			break
		}*/
		if cont := whetherContinue(); !cont {
			break
		}

	}
}

func whetherContinue() bool {
	var whetherContinue string
	fmt.Println("是否要录入下一个（y/n）?: ")
	fmt.Scanln(&whetherContinue)
	if whetherContinue != "y" {
		return false
	}
	return true
}

func mainFatRateBody() {
	age, weight, tall, sexWeight, sex := getMaterialsFromInput()

	fatRate := calcFatRate(weight, tall, age, sexWeight)

	getHealthnessSuggestions(sex, age, fatRate)
}

func getHealthnessSuggestions(sex string, age int, fatRate float64) {
	//建议
	if sex == "男" {
		//todo 男性体脂率建议表
		if age >= 18 && age <= 39 {
			if fatRate <= 0.1 {
				fmt.Println("目前是：偏瘦")
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				fmt.Println("目前是：正常")
			} else if fatRate > 0.16 && fatRate <= 0.26 {
				fmt.Println("目前是：偏胖")
			} else {
				fmt.Println("目前是：不健康")
			}
		} else if age >= 40 && age <= 59 {
			//todo
		} else if age >= 60 {
			//todo
		} else {
			fmt.Println("未成年人，无法评测")
			return
		}
	} else {
		//todo 女性体脂率建议表
	}
}

func calcFatRate(weight float64, tall float64, age int, sexWeight int) float64 {
	//jisuan
	var bmi float64 = weight / (tall * tall)
	var fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	fmt.Println("您的体脂率是：", fatRate)
	return fatRate
}

func getMaterialsFromInput() (int, float64, float64, int, string) {
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	var weight float64
	fmt.Print("体重（kg）：")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高（m）：")
	fmt.Scanln(&tall)

	var sexWeight int
	var sex string
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)
	if sex == "男" {
		sexWeight = 1
	} else if sex == "女" {
		sexWeight = 0
	} else {
		fmt.Println("输入有误")
	}
	return age, weight, tall, sexWeight, sex
}
