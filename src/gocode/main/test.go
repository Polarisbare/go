package main//声明文件所在包 每个go文件必须有所属包
import "fmt"//引入程序中需要用的包，为了使用包的函数 比如Println
func main()  {//main主函数 程序入口
	fmt.Println("Hello Golang")
}