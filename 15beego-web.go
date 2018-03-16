package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type HomeController struct {
	beego.Controller
}

func (this HomeController) Get() {
	this.Ctx.WriteString("Home.Get <script>(function(){alert('hello beege')})()</script>")

}

type CaseController struct {
	beego.Controller
}

func (this CaseController) Get() {
	this.Ctx.WriteString("Case.Get")
}

func main() {
	beego.Router("/home", &HomeController{})

	ns := beego.NewNamespace("/group",
		beego.NSGet("/user", func(ctx *context.Context) {
			ctx.WriteString("user")
		}),
		beego.NSRouter("/case", &CaseController{}),
		beego.NSGet("/issue", func(ctx *context.Context) {
			ctx.WriteString("issue")
		}),
	)
	beego.AddNamespace(ns)
	beego.Run(":9999")
}
