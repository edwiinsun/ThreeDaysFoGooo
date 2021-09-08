/*
 * @Author: @Edwin
 * @Date: 2021-09-08 20:02:49
 * @LastEditors: @Edwin
 * @LastEditTime: 2021-09-08 21:30:43
 * @Description: Go语言三天课程Day03    实战: 体脂计算器
 * @FilePath: \GoProject\src\Day03\main.go
 */
package main

import (
	"fmt"
	"math"
)

// 1. BMI= 体重(公斤) / (身高 * 身高) (米)
// 2. 体制率: 1.2 * BMI + 0.23   * 年龄 - 5.4-10.8 * 性别 (男1,女0)

func main() {
	for {

		//先声明变量
		var name, fatResult string
		var weight, tall, bmi, fatRate float64
		var age, sexFlag int
		var isMale bool

		//用户信息录入
		fmt.Print("姓名:")
		fmt.Scanln(&name)

		fmt.Print("性别(true:男性,false:女性):")
		fmt.Scanln(&isMale)
		//通过isMale来对性别flag进行赋值
		if isMale {
			sexFlag = 1
		} else {
			sexFlag = 0
		}

		fmt.Print("年龄:")
		fmt.Scanln(&age)
		//先对年龄进行判断, 及时终止程序
		if age < 18 {
			fmt.Println("你太小了, 别测了, 去学习吧!")
			break
		} else {
			fmt.Print("体重(KG):")
			fmt.Scanln(&weight)

			fmt.Print("身高(M):")
			fmt.Scanln(&tall)
		}

		//bmi计算方法
		bmi = weight / math.Pow(tall, 2)
		//体脂率计算方法
		fatRate = fatRate + (1.2*bmi + 0.23*float64(age) - 5.4 - 10*float64(sexFlag))
		// fmt.Println(fatRate)
		//结果
		if isMale {
			if age >= 18 && age <= 39 {
				//男性体脂率结果
				if fatRate <= 10 {
					fatResult = "偏瘦"
				} else if fatRate > 10 && fatRate <= 16 {
					fatResult = "标重"
				} else if fatRate > 16 && fatRate <= 21 {
					fatResult = "偏重"
				} else if fatRate > 21 && fatRate <= 26 {
					fatResult = "肥胖"
				} else {
					fatResult = "严重肥胖"
				}
			} else if age >= 40 && age <= 59 {
				if fatRate <= 11 {
					fatResult = "偏瘦"
				} else if fatRate > 11 && fatRate <= 17 {
					fatResult = "标重"
				} else if fatRate > 17 && fatRate <= 22 {
					fatResult = "偏重"
				} else if fatRate > 22 && fatRate <= 27 {
					fatResult = "肥胖"
				} else {
					fatResult = "严重肥胖"
				}

			} else {
				if fatRate <= 13 {
					fatResult = "偏瘦"
				} else if fatRate > 13 && fatRate <= 19 {
					fatResult = "标重"
				} else if fatRate > 19 && fatRate <= 24 {
					fatResult = "偏重"
				} else if fatRate > 24 && fatRate <= 29 {
					fatResult = "肥胖"
				} else {
					fatResult = "严重肥胖"
				}
			}

		} else {
			if age >= 18 && age <= 39 {
				//女性体脂率结果
				if fatRate <= 20 {
					fatResult = "偏瘦"
				} else if fatRate > 21 && fatRate <= 27 {
					fatResult = "标重"
				} else if fatRate > 27 && fatRate <= 34 {
					fatResult = "偏重"
				} else if fatRate > 34 && fatRate <= 39 {
					fatResult = "肥胖"
				} else {
					fatResult = "严重肥胖"
				}
			} else if age >= 40 && age <= 59 {
				if fatRate <= 21 {
					fatResult = "偏瘦"
				} else if fatRate > 21 && fatRate < 28 {
					fatResult = "标重"
				} else if fatRate > 28 && fatRate <= 35 {
					fatResult = "偏重"
				} else if fatRate > 35 && fatRate <= 40 {
					fatResult = "肥胖"
				} else {
					fatResult = "严重肥胖"
				}
			} else {
				if fatRate <= 22 {
					fatResult = "偏瘦"
				} else if fatRate > 22 && fatRate <= 29 {
					fatResult = "标重"
				} else if fatRate > 29 && fatRate <= 36 {
					fatResult = "偏重"
				} else if fatRate > 36 && fatRate <= 41 {
					fatResult = "肥胖"
				} else {
					fatResult = "严重肥胖"
				}
			}
		}

		fmt.Println(name, fatResult)

		whetherContinue := false
		fmt.Println("again?(true or false")
		fmt.Scanln(&whetherContinue)
		if !whetherContinue {
			break
		}
	}

}
