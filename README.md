# ThreeDaysFoGooo
三天Go语言简单入门  From 极客Live

📔Go语言学习记录

- Go语言的特点

  - Go语言对一些小众的编程哲学进行了吸收
    - 接受了函数式编程的一些想法, 支持匿名函数与闭包
    - 接受了以Erlang语言为代表的面向消息编程思想, 支持goroutine和通道
  - 主要特性🎈
    - 自动垃圾回收
    - 更丰富的内置类型
    - 函数多返回值
    - 错误处理
    - 匿名函数和闭包
    - 类型和接口
    - 并发编程
    - 反射
    - 语言交互性

- 学习里程🏁

  - 第一天 

    - Go安装路径: D:\Golang\

    - 环境变量配置

      - 用户环境变量新建
        - 变量名: GOROOT  变量值:  D:\Golang\
        - 变量名: GOPATH  变量值:  D:\GoProject\
          - 目录下新增三个文件夹: bin src pkg
        - Path环境变量新增
          - %GOROOT%\bin
          - %GOPATH%\bin

    - Go语言初体验

      package main import "fmt" *// or* *// import (* *//  "fmt"* *// )* func main() { *// 第一个括号不可以换到下一行*  fmt.Println("Hello World!") } 

    - 

- 第二天

  - 变量(强类型数据语言)  变量有严格的数据类型限制    *//可以用utf-8里的语言当作变量的名字*

    - 数据类型

      var hello string = "Hello World!"      // 注意语法 arr:=[]int{2,3,4,5}   //数组的定义

      - int、float64、struct、map

    - 如何定义

      - 有很多种定义变量的方式

    - 如何使用

    - 变量默认值

      - 在所有数字类型中,默认值都是0
      - 字符串定义过程中, 默认值是空行(不是null)
      - map是key --- value键值对, 默认值是map[](也是空的)

    - 变量类型推断

      - 定义变量过程中可以不指定变量类型， golang会自行推断

        var v3,v4="c",4  fmt.Println(v3,v4)  fmt.Println(reflect.TypeOf(v3))  fmt.Println(reflect.TypeOf(v4))

    - 变量类型转换及其代价

      - ![img](https://api2.mubu.com/v3/document_image/bb1b0305-df55-43e6-982e-5f0d2c688eb1-13872488.jpg)
      - ![img](https://api2.mubu.com/v3/document_image/ac1a13bb-192b-4edf-8fdd-dc5e3963f126-13872488.jpg)

      - 数字可以互相转化， 但是有代价 ： 浮点型转换为int会丢失所有小数，从一个无符号整数转换成有符号整数

    - 变量使用法则

      - \1. Golang是一个强类型语言，前后赋值类型必须一致
      - \2. 变量必须先定义
      - \3. Golang会为每个变量设置默认值
      - \4. 在同一个作用域内变量不可以重名

  - 常量

    - 常量与变量不同的地方
      - \1. 变量不调用不会提示报错
      - \2. 常量是在编译过程中计算好的值
      - \3. 常量只能支持: bool  string  数字
      - \4. 变量是在运行过程中配置的值
      - \5. 变量没有限制

  - 计算符号

    - 常规运算： 加、减、乘、除、取余        

      \* 当两个无符号变量相减时， 结果有问题 *取余只支持int

    - 比较运算符： &&、||、!、==、!=、> 、<、>=、<=

    - 位运算符： &、|、<<、>>、^异或       

      - ![img](https://api2.mubu.com/v3/document_image/9b6e5b43-bede-41be-9220-8316afc39b7b-13872488.jpg)

      移位对于乘除运算效率极高

    - 条件表达式： if-else

    - 循环：for

  - 字符串及其操作

    - 字符串运算    : 相加

    - 格式化输出     : 具体查看Golang官方文档

      Printf: 格式化输出,   fmt.Printf("%s \t %d \n \\  %%", str1, i)

      - %.2f  精确两位小数
      - %s   string
      - %d  整数

- 第三天

  - 体脂计算器
    - 题目要求: BMI计算法
      - \1. BMI= 体重(公斤) / (身高 * 身高) (米)
      - \2. 体制率: 1.2 * BMI / 0.23   * 年龄 - 5.4-10.8 * 性别 (男1,女0)
    - 计算体脂率, 并告诉偏瘦、标准、偏重、肥胖、严重肥胖
