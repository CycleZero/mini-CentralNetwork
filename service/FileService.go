package service

import (
	"encoding/json"

	constformat "../ConstFormat"
)

type FileService struct {
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
func (f *FileService) HundleCommand(command constformat.NetCommandPackage) {
	comobj := FileCommand{}
	err := json.Unmarshal([]byte(command.Commandpackage.Command), &comobj)
	if err != nil {
		return
	}
}

func (f *FileService) Init() {

}
