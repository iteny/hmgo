package admin

import (
	// "fmt"
	// "github.com/iteny/hmgo"
	// "net/http"2223231111123123222sdfs
	// "github.com/go-ini/ini"
	// "fmt"
	"github.com/kataras/iris"
	"hmgo/conf"
)

type Login struct{}

func (l *Login) Index(ctx *iris.Context) {
	site := conf.Cfg.Section("siteset").Key("WEB_SITE_TITLE")
	ctx.Render("admin/login/hi.html", map[string]interface{}{"meta_title": "后台登录", "site_title": site}, iris.RenderOptions{"gzip": true})
}
