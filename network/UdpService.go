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
	InChan   chan constformat.UDPdata
}

func (u *UdpService) Init(host string, port int) {
	u.addr = &net.UDPAddr{IP: net.ParseIP(host), Port: port}
	u.OutChan = make(chan constformat.UDPdata, 100)
	u.InChan = make(chan constformat.UDPdata, 100)
	u.isInit = true
}

func (u *UdpService) Run() {
	if !u.isInit {
		return
	}
	listen, err := net.ListenUDP("udp", u.addr)
	if err != nil {
		fmt.Println("Listen failed, err: ", err)
		return
	}
	u.listener = listen
	go u.Reader()
	go u.Sender()
	fmt.Println("UDP服务已启动,监听端口:", u.addr.Port)
}

func (u *UdpService) Reader() {
	var data [10240]byte

	for {
		n, addr, err := u.listener.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed, err: ", err)
			break
		}

		fmt.Println("收到来自", addr, "的数据：", string(data[:n]))
		u.OutChan <- u.GenerateUDPdata(data[:n], addr)
		fmt.Println("数据已交由DataBridge处理")
	}
}

func (u *UdpService) Sender() {
	defer u.listener.Close()
	for {
		data := <-u.InChan
		_, err := u.listener.WriteToUDP(data.Data, data.ToAddr)
		if err != nil {
			fmt.Println("send error:", err)
			break
		}
	}
}

func (u *UdpService) GenerateUDPdata(data []byte, addr *net.UDPAddr) constformat.UDPdata {
	return constformat.UDPdata{Data: data, FromAddr: addr, ToAddr: u.addr}
}
