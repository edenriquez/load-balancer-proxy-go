package server

import (
	"github.com/edenriquez/load-balancer-proxy-go/api/handlers"
	service "github.com/edenriquez/load-balancer-proxy-go/pkg/proxy"
	"github.com/kataras/iris"
)

// RouterSetUp shoulddeclare routes
func RouterSetUp(app *iris.Application, service *service.Service) {
	app.Post("/", func(c iris.Context) {
		handlers.ProxyHandler(c, service)
	})
}
