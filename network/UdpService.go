package network

import (
	"fmt"
	"net"

	constformat "../ConstFormat"
)

type UdpService struct {
	addr     *net.UDPAddr
	isInit   bool
	listener *net.UDPConn
	OutChan  chan constformat.UDPdata
}

func (u *UdpService) Init(host string, port int) {
	u.addr = &net.UDPAddr{IP: net.ParseIP(host), Port: port}
	u.OutChan = make(chan constformat.UDPdata, 100)
	u.isInit = true
}

func (u *UdpService) Run() {
	if u.isInit == false {
		return
	}
	listen, err := net.ListenUDP("udp", u.addr)
	if err != nil {
		fmt.Println("Listen failed, err: ", err)
		return
	}
	u.listener = listen
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("read udp failed, err: ", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil {
			fmt.Println("Write to udp failed, err: ", err)
			continue
		}
	}
}

func (u *UdpService) Reader() {
	var data [10240]byte
	for {
		n, addr, err := u.listener.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed, err: ", err)
			break
		}
		u.OutChan <- u.GenerateUDPdata(data[:n], addr)
	}
}

func (u *UdpService) GenerateUDPdata(data []byte, addr *net.UDPAddr) constformat.UDPdata {
	return constformat.UDPdata{Data: data, FromAddr: addr, ToAddr: u.addr}
}
