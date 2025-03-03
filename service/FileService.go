package service

import (
	"encoding/json"
	"os"

	constformat "../ConstFormat"
)

type FileService struct {
	CurrentFilePath string
	BasePath        string
}

type FileCommand struct {
	con      string // list download upload delete preview
	filename string
	filepath string
	fileid   int
}

type FileObject struct {
	name  string
	isDir bool
	size  int
	path  string
}

type FileServiceResponse struct {
	data any
	id   int
	code string
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
func (f *FileService) HundleCommand(command constformat.NetCommandPackage) FileServiceResponse {
	comobj := FileCommand{}
	err := json.Unmarshal([]byte(command.Commandpackage.Command), &comobj)
	if err != nil {
		return FileServiceResponse{code: "json parse error", data: ""}
	}
	response := FileServiceResponse{}
	switch comobj.con {
	case "delete":
		_, err := f.DeleteFile(comobj.filepath, comobj.filename)
		if err != nil {
			// return FileServiceResponse{code: "delete file error", data: ""}
			response.code = "delete file error"
			response.data = err.Error()
			return response
		}
		// return FileServiceResponse{code: "delete file success", data: comobj.filepath + "," + comobj.filename}
		response.code = "delete file success"
		response.data = comobj.filepath + "," + comobj.filename
	case "list":
		filelist := ListFilesAndDirs(comobj.filepath)
		// return FileServiceResponse{code: "list file success", data: filelist}
		response.code = "list file success"
		response.data = filelist
		// 处理 list 命令
	case "download":
		// 处理 download 命令
		// return FileServiceResponse{code: "download file success", data: ""}
		response.code = "download file success"
		// response.data = ""
	case "upload":
		// 处理 upload 命令
		// return FileServiceResponse{code: "upload file success", data: ""}
	case "preview":
		// 处理 preview 命令
		// return FileServiceResponse{code: "preview file success", data: ""}
	default:
		// 处理未知命令
		// return FileServiceResponse{code: "unknown command", data: ""}
		response.code = "unknown command"
		response.data = ""
	}
	return response

}

func (f *FileService) New(basePath string) *FileService {
	return &FileService{BasePath: basePath}
}

func ListFilesAndDirs(path string) []FileObject {
	if path == "" {
		drives := ListDiskDrives()
		filelist := []FileObject{}
		for _, drive := range drives {
			filelist = append(filelist, FileObject{name: drive, isDir: true, path: drive})
		}
		return filelist
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil
	}
	filelist := []FileObject{}
	for _, entry := range entries {
		fileobj := FileObject{name: entry.Name(), path: path + "/" + entry.Name()}
		if entry.IsDir() {
			fileobj.isDir = true
		} else {
			fileobj.isDir = false
		}
		filelist = append(filelist, fileobj)
	}
	return filelist
}

func (f *FileService) DeleteFile(filepath, filename string) (bool, error) {
	fullPath := filepath + "/" + filename
	err := os.Remove(fullPath)
	if err != nil {
		return false, err
	}
	return true, nil
}

func FileObjectToString(file FileObject) string {
	b, err := json.Marshal(file)
	if err != nil {
		return ""
	}
	return string(b)
}

func FileListToString(filelist []FileObject) string {
	s := ""
	for _, file := range filelist {
		s += FileObjectToString(file) + ","
	}
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	s = "[" + s + "]"
	return s
}

// 列出所有可用盘符
func ListDiskDrives() []string {
	drives := make([]string, 0)
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		path := string(drive) + ":"
		if _, err := os.Stat(path); err == nil {
			drives = append(drives, path)
		}
	}
	return drives
}
