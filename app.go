package hmgo

import (
	"fmt"
	"net/http"
)

var HmApp *App

func init() {
	HmApp = NewApp()
}

type App struct {
	Handlers *ControllerRegister
	Server   *http.Server
}

func NewApp() *App {
	cr := NewControllerRegister()
	app := &App{Handlers: cr, Server: &http.Server{}}
	return app
}

//运行web服务器
func (a *App) Run() {
	addr := HmConfig.Listen.HttpAddr
	if HmConfig.Listen.HttpPort != 0 {
		addr = fmt.Sprintf("%s:%d", HmConfig.Listen.HttpAddr, HmConfig.Listen.HTTPSPort)
	}
}
