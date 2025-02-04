package bridge

import (
	"encoding/json"
	"fmt"
	"net"

	constformat "../ConstFormat"
	"../service"
)

type DataBridge struct {
	Filecontainer    *service.FileContainer
	Messagecontainer *service.MessageContainer

	ServiceOutchan chan constformat.ServiceData
	UDPInchan      chan constformat.UDPdata
	UDPOutchan     chan constformat.UDPdata
}

func (d *DataBridge) TransCommand(pkg constformat.NetCommandPackage) {
	switch {
	case pkg.Commandpackage.TargetService == "FileContainer":
		d.Filecontainer.HundleCommand(pkg)
	case pkg.Commandpackage.TargetService == "MessageContainer":
		d.Messagecontainer.InChan <- pkg
		fmt.Println("数据已送达MessageContainer")
	}

}

func (d *DataBridge) Init() {
	d.ServiceOutchan = make(chan constformat.ServiceData, 100)

	d.Filecontainer = new(service.FileContainer)
	d.Filecontainer.Init()
	d.Messagecontainer = new(service.MessageContainer)
	d.Messagecontainer.Init(d.ServiceOutchan)

}

func (d *DataBridge) Run() {
	d.Messagecontainer.Run()

	go d.runServiceOut()
	go d.runUDPOut()
	fmt.Println("DataBridge 服务已启动")
}

func (d *DataBridge) DecodeData(data constformat.UDPdata) constformat.NetCommandPackage {
	jsonData := constformat.CommandPackage{}
	err := json.Unmarshal(data.Data, &jsonData)
	if err != nil {
		fmt.Println("json解析错误", err)
	}
	result := constformat.NetCommandPackage{Commandpackage: jsonData, ToAddr: &net.UDPAddr{IP: data.FromAddr.IP, Port: 45689}}
	fmt.Println("解析得到NetCommandPackage:", result)
	return result

}

func (d *DataBridge) runServiceOut() {
	fmt.Println("ServiceOut-DataBridge-UDPin通道已接通")
	for {
		select {
		case data := <-d.ServiceOutchan:
			fmt.Println("接收到来自ServiceOut的数据")
			d.UDPInchan <- constformat.UDPdata{Data: data.Data, ToAddr: data.ToAddr, Id: data.Id}
			fmt.Println("数据已送达UDPInchan")
		}
	}
}

func (d *DataBridge) runUDPOut() {
	fmt.Println("DataBridge-UDPout通道已接通")
	for {
		select {
		case data := <-d.UDPOutchan:
			fmt.Println("接收到来自UDPOut的数据")
			d.TransCommand(d.DecodeData(data))
		}
	}
}
