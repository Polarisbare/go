/*
 * @Author: Lv Jingxin lv510987@163.com
 * @Date: 2023-05-25 17:31:04
 * @LastEditors: Lv Jingxin lv510987@163.com
 * @LastEditTime: 2023-06-09 17:32:00
 * @FilePath: /gocode/unit2/demo2/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
func main(){
    // & ： 返回变量的储存地址
    // * 取指针变量对应的数值
    var age int = 18
    fmt.Println("age的内存地址",&age)
    var ptr *int = &age
    fmt.Println(ptr)
    fmt.Println("ptr的内存地址",&ptr)
}