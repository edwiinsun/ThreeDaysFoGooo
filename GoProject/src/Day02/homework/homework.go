/*
 * @Author: @Edwin
 * @Date: 2021-09-07 21:50:11
 * @LastEditors: @Edwin
 * @LastEditTime: 2021-09-07 22:24:12
 * @Description: 两个圆的面积之和, 结果保留三位小数
 * @FilePath: \GoProject\src\Day02\homework\homework.go
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = 3.14
	var radius1, radius2 float64 = 3, 4
	var area1 float64 = pi * (math.Pow(radius1, 2))
	var area2 float64 = pi * (math.Pow(radius2, 2))

	fmt.Printf("%.3f", area1+area2)

}
