package service

import constformat "../ConstFormat"

type MessageContainer struct {
	OutChan chan constformat.ServiceData
	InChan  chan constformat.NetCommandPackage
}

func (m *MessageContainer) GetMessage() string {
	return "Hello World"
}

func (m *MessageContainer) HundleCommand(pkg constformat.NetCommandPackage) {
	switch {
	case pkg.Command == "GetMessage":
		m.OutChan <- constformat.ServiceData{Data: []byte(m.GetMessage())}
	}
}

func (m *MessageContainer) Init() {
	m.InChan = make(chan constformat.NetCommandPackage, 100)
	m.OutChan = make(chan constformat.ServiceData, 100)
}

func (m *MessageContainer) Run() {
	for {
		select {
		case pkg := <-m.InChan:
			m.HundleCommand(pkg)
		}

	}
}
