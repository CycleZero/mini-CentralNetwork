package service

import (
	"fmt"

	constformat "../ConstFormat"
)

type MessageContainer struct {
	OutChan chan constformat.ServiceData
	InChan  chan constformat.NetCommandPackage
}

func (m *MessageContainer) GetMessage() string {
	return "Hello World"
}

func (m *MessageContainer) HundleCommand(pkg constformat.NetCommandPackage) {

	m.OutChan <- constformat.ServiceData{Data: []byte(m.GetMessage()), ToAddr: pkg.ToAddr, Id: pkg.Commandpackage.Id}
	fmt.Println("MessageContainer 服务已处理", pkg.ToAddr)
}

func (m *MessageContainer) Init(outchan chan constformat.ServiceData) {
	m.InChan = make(chan constformat.NetCommandPackage, 100)
	m.OutChan = outchan
}

func (m *MessageContainer) Run() {

	go func() {
		for {
			select {
			case pkg := <-m.InChan:
				m.HundleCommand(pkg)
			}

		}
	}()
	fmt.Println("MessageContainer 服务已启动")
}
