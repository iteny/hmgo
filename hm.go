package hmgo

import (
	// "os"
	"strconv"
	"strings"
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
	if len(params) > 0 && params[0] != "" {
		strs := strings.Split(params[0], ":")
		if len(strs) > 0 && strs[0] != "" {
			HmConfig.Listen.HttpAddr = strs[0]
		}
		if len(strs) > 1 && strs[1] != "" {
			HmConfig.Listen.HttpPort, _ = strconv.Atoi(strs[1])
		}
	}
	HmApp.Run()
}

//运行http服务前加载hook
func loadBeforeHttpRun() {
	addHook(loadMine)
	addHook(loadHttpErrorHandler)
	// addHook(loadSession)
	// addHook(loadTemplate)
	// addHook(loadAdmin)
	// addHook(loadGzip)
	for _, hk := range hooks {
		if err := hk(); err != nil {
			panic(err)
		}
	}
}
