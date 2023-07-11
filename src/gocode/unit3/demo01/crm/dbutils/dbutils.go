package dbutils

import "fmt"

func getConn(){//首字母大写可以被其他包访问，小写不可以
	fmt.Println("执行啦dbutils下的getConn函数")
}