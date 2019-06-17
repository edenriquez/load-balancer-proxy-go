package server

import (
	"github.com/edenriquez/load-balancer-proxy-go/api/handlers"
	"github.com/edenriquez/load-balancer-proxy-go/api/middlewares"
	service "github.com/edenriquez/load-balancer-proxy-go/pkg/proxy"
	"github.com/kataras/iris"
)

// RouterSetUp shoulddeclare routes
func RouterSetUp(app *iris.Application, service *service.Service) {
	app.Get("/", func(c iris.Context) {
		middlewares.Scheduler(c, service)
	}, handlers.ProxyHandler)
}
