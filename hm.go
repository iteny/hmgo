package main

import (
	"github.com/kataras/iris"
	_ "hmgo/conf"
	_ "hmgo/routers"
)

func main() {

	iris.Listen("sina.com:80")
}
