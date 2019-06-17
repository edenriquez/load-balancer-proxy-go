package handlers

import (
	"fmt"

	service "github.com/edenriquez/load-balancer-proxy-go/pkg/proxy"
	"github.com/kataras/iris"
)

// ProxyHandler should redirect to specific domain
func ProxyHandler(ctx iris.Context, service *service.Service) {
	domain := ctx.GetHeader("domain")
	result, err := service.Find(domain)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(iris.Map{"result": "not found"})
		return
	}
	ctx.JSON(iris.Map{"result": result})
}
