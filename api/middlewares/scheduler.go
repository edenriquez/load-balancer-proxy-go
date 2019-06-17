package middlewares

import (
	"fmt"

	"github.com/edenriquez/load-balancer-proxy-go/pkg/entity"
	service "github.com/edenriquez/load-balancer-proxy-go/pkg/proxy"
	"github.com/kataras/iris"
)

// proxyQueue is the request queue
var proxyQueue []*entity.Proxy

// StartQueue should initialize proxy queue
func StartQueue() {
	proxyQueue = append(proxyQueue, &entity.Proxy{})
}

// Scheduler should store queue requests
func Scheduler(ctx iris.Context, service *service.Service) {
	domain := ctx.GetHeader("domain")
	result, err := service.Find(domain)
	if err != nil {
		ctx.JSON(iris.Map{"result": err})
		return
	}
	queue(result)
	ctx.Next()
}

/*
 Queue should queue based in priority and weight properties
	example of swap values
	proxyQueue := [1,2,3]
	p := 4
	i := 1
	copy(proxyQueue[i+1:], proxyQueue[i:])
	[3] => [2,3]


	proxyQueue[i] = p
	[1,(2),2,3] = [4]
	[1,4,2,3]
*/
func queue(p *entity.Proxy) {
	fmt.Println("LEN ANTES", len(proxyQueue))
	if len(proxyQueue) == 0 {
		proxyQueue = append(proxyQueue, p)
		return
	}
	for i, proxy := range proxyQueue {
		if proxy.Weight+proxy.Priority > p.Weight*p.Priority {
			proxyQueue = append(proxyQueue, &entity.Proxy{})

			copy(proxyQueue[i+1:], proxyQueue[i:])

			proxyQueue[i] = p
		}
	}
}

// Dequeue should de-queue last item
func Dequeue() *entity.Proxy {
	var result *entity.Proxy
	for len(proxyQueue) > 0 {
		result = proxyQueue[0]
		proxyQueue = proxyQueue[1:]
	}
	return result
}
