package main

import (
	"blog/common/logfile"
	"blog/common/validate"
	"fmt"
)

func main() {
	/*logfile.Info("测试")
	router.LoadRouter()
	route.Run(4567)*/
	initServer()

}

func initServer() {
	// 初始化认证器
	if err := validate.InitValidator(); err != nil {
		fmt.Println("load validator error: ", err)
		panic(err)
	}
	logfile.InitLogfile()
}
