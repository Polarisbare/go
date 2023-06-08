/*
 * @Author: Lv Jingxin lv510987@163.com
 * @Date: 2023-05-25 15:59:07
 * @LastEditors: Lv Jingxin lv510987@163.com
 * @LastEditTime: 2023-05-25 16:10:40
 * @FilePath: /gocode/unit2/demo1/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main(){
	// int 整数
	var age int
	age = 18
	fmt.Println("age=",age)

	var age2 int = 19
	fmt.Println("age2=",age2)

	// 制定变量的类型 不赋值 使用默认值
	var num2 int
	fmt.Println(num2)

	// 第三种 如果没有写变量的类型，那么根据后面的值进行判定变量的类型 自动类型推断
	var num3 = "tom"
	fmt.Println(num3)

	// 第四种 省略var 注意 := 不能写为 =
	sex := "男"
	fmt.Println(sex)
}