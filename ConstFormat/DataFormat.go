package constformat

import "net"

type NetCommandPackage struct {
	Commandpackage CommandPackage
	ToAddr         *net.UDPAddr
}

type CommandPackage struct {
	Command       string
	Id            string
	TargetService string
}

type UDPdata struct {
	Data     []byte
	FromAddr *net.UDPAddr
	ToAddr   *net.UDPAddr
	Id       string
}

type ServiceData struct {
	Data   []byte
	ToAddr *net.UDPAddr
	Id     string
}
