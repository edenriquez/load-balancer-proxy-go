package main

import (
	"github.com/edenriquez/load-balancer-proxy-go/api/server"
	"github.com/edenriquez/load-balancer-proxy-go/api/utils"
	service "github.com/edenriquez/load-balancer-proxy-go/pkg/proxy"
	repository "github.com/edenriquez/load-balancer-proxy-go/pkg/repository"
)

func main() {
	utils.LoadEnvVars()

	proxyConn := repository.NewMysqlRepository()
	proxyService := service.NewService(proxyConn)
	proxyService.Migrate()

	settings := server.SetUp()
	server.RouterSetUp(settings, proxyService)
	server.RunServer(settings)
}
