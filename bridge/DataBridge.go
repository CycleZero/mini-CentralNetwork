package bridge

import (
	constformat "../ConstFormat"
	"../service"
)

type DataBridge struct {
	Filecontainer *service.FileContainer
}

func (d *DataBridge) TransCommand(pkg constformat.NetCommandPackage) {
	switch {
	case pkg.TargetService == "FileContainer":
		d.Filecontainer.HundleCommand(pkg)
	}

}

func (d *DataBridge) Init() {
	d.Filecontainer = new(service.FileContainer)

}
