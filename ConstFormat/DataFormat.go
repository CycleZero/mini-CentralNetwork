package constformat

import "net"

type NetCommandPackage struct {
	Command       string
	ResultTarget  chan string
	Conn          net.Conn
	Id            string
	TargetService string
}

type UDPdata struct {
	Data     []byte
	FromAddr *net.UDPAddr
	ToAddr   *net.UDPAddr
	Id       string
}
