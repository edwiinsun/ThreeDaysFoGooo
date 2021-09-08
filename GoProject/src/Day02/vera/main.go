/*
 * @Author: @Edwin
 * @Date: 2021-09-07 20:11:56
 * @LastEditors: @Edwin
 * @LastEditTime: 2021-09-07 21:29:11
 * @Description: &1
 * @FilePath: \GoProject\src\Day02\vera\main.go
 */

package main

import (
	"fmt"
	"math"
	"reflect"
)

// or

// import (
// 	"fmt"
// )

func main() {
	var hello string = "Hello World!"
	//创建hello变量， 注意语法
	fmt.Println(hello)
	var int2 int = 33
	fmt.Println(int2)
	var float2 float64 = 1.234
	fmt.Println(float2)

	var struct4 struct{} = struct{}{}
	fmt.Println(struct4)
	var map5 map[string]string = map[string]string{}
	fmt.Println(map5)

	var hello1, hello2 string = "nihao ", "golang"
	//Go语言的变量可以成串定义
	fmt.Println(hello1, hello2)
	var (
		hello4 = "nihao"
		hello3 = "golang"
	) //方法快定义
	fmt.Println(hello4, hello3)

	//简洁表示
	studentName := "sjh"
	fmt.Println(studentName)

	//可以用utf-8里的语言当作变量的名字!!
	横 := 3
	竖 := 4
	面积 := 横 * 竖
	fmt.Println(面积)

	var v3, v4 = "c", 4
	fmt.Println(v3, v4)
	fmt.Println(reflect.TypeOf(v3))
	fmt.Println(reflect.TypeOf(v4))

	//变量转化及其代价
	var f1 float64 = 1.23456
	var i1 int64 = int64(f1)
	fmt.Println(f1)
	fmt.Println(i1)

	var ui1 uint64 = math.MaxUint64
	var i2 int64 = int64(ui1)
	fmt.Println(ui1)
	fmt.Println(i2)

	//常量
	const pi float64 = 3.14

	//计算符号
	var left, right int64 = 2, 3
	value := left + right
	fmt.Println(value)

	//条件表达式
	// 	func main2() {
	// 		var buyFruit string ="6个苹果"
	// 		var watermellon bool =false
	// 		if watermellon{
	// buyFruit="1 个苹果"
	// 		}
	// 		fmt.Println(buyFruit)
	// 	}

	//循环
	// func main(){
	// 	for i:=0;i<100;i++{
	// 		fmt.Println("...")
	// 	}
	// }
	//简便写法
	// func main(){
	//i:=0;
	// 	for i<100{
	// 		fmt.Println("...")
	// 	}
	// }
}
