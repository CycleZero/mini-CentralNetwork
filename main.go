package main

import (
	"fmt"

	"./bridge"
	"./network"
)

func main() {
	run()

}

func udptest() {
	udpobj := network.UdpService{}
	udpobj.Init("127.0.0.1", 45688)
	udpobj.Run()

}

func run() {
	databridge := &bridge.DataBridge{}
	databridge.Init()
	udp := &network.UdpService{}
	udp.Init("127.0.0.1", 45688)
	databridge.UDPInchan = udp.InChan
	databridge.UDPOutchan = udp.OutChan
	udp.Run()
	databridge.Run()
	fmt.Println("所有服务已启动 ")
	for {

	}
}
