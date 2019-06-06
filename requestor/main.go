package main

import serial "github.com/edenriquez/load-balancer-proxy-go/requestor/serial"
import flagger "github.com/edenriquez/load-balancer-proxy-go/requestor/flagger"

func main() {
	args := flagger.Flagger()
	serial.Requestor(args...)
}
