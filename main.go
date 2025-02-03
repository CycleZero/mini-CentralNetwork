package main

import (
	"./network"
)

func main() {
	udptest()

}

func udptest() {
	udpobj := network.UdpService{}
	udpobj.Init("127.0.0.1", 45688)
	udpobj.Run()

}
