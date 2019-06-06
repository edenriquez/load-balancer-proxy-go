package main

import (
	"github.com/edenriquez/load-balancer-proxy-go/requestor/concurrent"
	flagger "github.com/edenriquez/load-balancer-proxy-go/requestor/flagger"
)

func main() {
	args := flagger.Flagger()
	// serial.Requestor(args...)
	concurrent.Requestor(args...)
}
