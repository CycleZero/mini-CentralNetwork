package bridge

import (
	constformat "../ConstFormat"
	"../service"
)

type DataBridge struct {
	Filecontainer    *service.FileContainer
	Messagecontainer *service.MessageContainer
}

func (d *DataBridge) TransCommand(pkg constformat.NetCommandPackage) {
	switch {
	case pkg.TargetService == "FileContainer":
		d.Filecontainer.HundleCommand(pkg)
	}

}

func (d *DataBridge) Init() {
	d.Filecontainer = new(service.FileContainer)
	d.Filecontainer.Init()
	d.Messagecontainer = new(service.MessageContainer)
	d.Messagecontainer.Init()
}

func (d *DataBridge) Run() {
	d.Messagecontainer.Run()

	for {
		select {}

	}
}
