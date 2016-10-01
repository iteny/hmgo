package hmgo

import (
	"os"
)

const (
	VERSION = "0.1"
	DEV     = "dev"
)

var hooks = make([]hookFunc, 0)

//定义hookFunc的类型
type hookFunc func() error

//添加hook
func addHook(hf hookFunc) {
	hooks = append(hooks, hf)
}

/**
 * hmgo.Run() 运行默认的addrPort
 * hmgo.Run("localhost")
 * hmgo.Run(":8088")
 * hmgo.Run("127.0.0.1:8088")
 */
func Run(params ...string) {
	loadBeforeHttpRun()
}

//运行http服务前加载hook
func loadBeforeHttpRun() {
	addHook()
}
