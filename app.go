package hmgo

import (
	"fmt"
	// "log"
	// "net"
	"net/http"
	// "time"
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
	// mux := &RouteMux{}
	// mux := HttpNew()
	// mux.Get("/:last/:first", Whoami)
	fmt.Printf("%#v", &RouteMux{})

	http.ListenAndServe(":8088", nil)
	// var (
	// 	err        error
	// 	l          net.Listener
	// 	endRunning = make(chan bool, 1)
	// )
	// addr := HmConfig.Listen.HttpAddr
	// if HmConfig.Listen.HttpPort != 0 {
	// 	addr = fmt.Sprintf("%s:%d", HmConfig.Listen.HttpAddr, HmConfig.Listen.HTTPSPort)
	// }
	// a.Server.Handler = a.Handlers
	// a.Server.ReadTimeout = time.Duration(HmConfig.Listen.ServerTimeOut) * time.Second
	// a.Server.WriteTimeout = time.Duration(HmConfig.Listen.ServerTimeOut) * time.Second
	// // a.Server.ErrorLog = logs.GetLogger("HTTP")
	// if HmConfig.Listen.EnableHTTP {
	// 	go func() {
	// 		app.Server.Addr = addr
	// 		logs.Info("http server Running on http://%s", app.Server.Addr)
	// 		if BConfig.Listen.ListenTCP4 {
	// 			ln, err := net.Listen("tcp4", app.Server.Addr)
	// 			if err != nil {
	// 				logs.Critical("ListenAndServe: ", err)
	// 				time.Sleep(100 * time.Microsecond)
	// 				endRunning <- true
	// 				return
	// 			}
	// 			if err = app.Server.Serve(ln); err != nil {
	// 				logs.Critical("ListenAndServe: ", err)
	// 				time.Sleep(100 * time.Microsecond)
	// 				endRunning <- true
	// 				return
	// 			}
	// 		} else {
	// 			if err := app.Server.ListenAndServe(); err != nil {
	// 				logs.Critical("ListenAndServe: ", err)
	// 				time.Sleep(100 * time.Microsecond)
	// 				endRunning <- true
	// 			}
	// 		}
	// 	}()
	// }
	// <-endRunning
}
