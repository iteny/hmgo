package routers

import (
	// "fmt"
	"github.com/kataras/go-template/django"
	"github.com/kataras/iris"
	"hmgo/cons/admin"
	// "net/http"11111
)

func init() {

	iris.UseTemplate(django.New()).Directory("./templates", ".html")
	admingroup := iris.Party("admin.")
	{
		login := &admin.Login{}
		admingroup.Get("/", login.Index)
		admingroup.Static("/static", "./static/", 1)
		admingroup.Get("/something", dynamicSubdomainHandler)
		admingroup.Get("/something/:param1", dynamicSubdomainHandlerWithParam)
	}

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hello from mydomain.com path: %s", ctx.PathString())
	})

	iris.Get("/hello", func(ctx *iris.Context) {
		ctx.Write("Hello from mydomain.com path: %s", ctx.PathString())
	})
}
func dynamicSubdomainHandler(ctx *iris.Context) {
	username := ctx.Subdomain()
	ctx.Write("Hello from dynamic subdomain path: %s, here you can handle the route for dynamic subdomains, handle the user: %s", ctx.PathString(), username)
	// if  http://username4.mydomain.com:8080/ prints:
	// Hello from dynamic subdomain path: /, here you can handle the route for dynamic subdomains, handle the user: username4
}

func dynamicSubdomainHandlerWithParam(ctx *iris.Context) {
	username := ctx.Subdomain()
	ctx.Write("Hello from dynamic subdomain path: %s, here you can handle the route for dynamic subdomains, handle the user: %s", ctx.PathString(), username)
	ctx.Write("THE PARAM1 is: %s", ctx.Param("param1"))
}
