package handlers

import (
	scheduler "github.com/edenriquez/load-balancer-proxy-go/api/middlewares"
	"github.com/kataras/iris"
)

// ProxyHandler should redirect to specific domain
func ProxyHandler(ctx iris.Context) {
	lastIn := scheduler.Dequeue()
	ctx.JSON(iris.Map{"result": lastIn})
}
