package main

import (
	
	"net/http"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this HomeController) Get() {
	this.Ctx.WriteString("Home.Get <script>(function(){alert('hello beege')})()</script>")
	
}

func main() {
	beego.Router("/home", &HomeController{})
	beego.NSGet("/user",func(ctx *context.Context){
		ctx.
	})
	beego.Run()
}
