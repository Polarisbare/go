/*
 * @Author: Lv Jingxin lv510987@163.com
 * @Date: 2023-05-25 17:31:04
 * @LastEditors: Lv Jingxin lv510987@163.com
 * @LastEditTime: 2023-05-29 10:38:43
 * @FilePath: /gocode/unit2/demo2/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
func main(){
	// var stockcode=123
    // var enddate="2020-12-31"
    // var url="Code=%d&endDate=%s"
    // fmt.Println(5==9 || 6<7)
    // 没有三等 只有两等 
    var username string
    fmt.Println("请输入用户名：")
    fmt.Scanf("%s",&username)
    if(username == "无敌"){
        fmt.Println("登陆成功")
    }else{
        fmt.Println("登陆失败")
    }

}