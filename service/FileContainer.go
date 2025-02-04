package service

import (
	"encoding/json"

	constformat "../ConstFormat"
)

type FileContainer struct {
	FilePath string
	FileSize int64
}

type FileCommand struct {
	con      string
	filename string
	filepath string
	fileid   int
}

/*
command

	{
		con string
		filename string
		filepath string
		[fileid] int

}
*/
func (f *FileContainer) HundleCommand(command constformat.NetCommandPackage) {
	comobj := FileCommand{}
	err := json.Unmarshal([]byte(command.Commandpackage.Command), &comobj)
	if err != nil {
		return
	}
}

func (f *FileContainer) Init() {

}
