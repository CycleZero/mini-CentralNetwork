package service

import (
	"encoding/json"
	"os"
	"time"

	constformat "../ConstFormat"
)

type FileService struct {
	CurrentFilePath string
	BasePath        string
}

type FileCommand struct {
	Con      string // list download upload delete preview
	Filename string
	Filepath string
	FileId   int
}

type FileObject struct {
	Name    string
	IsDir   bool
	Size    int64
	Path    string
	ModTime time.Time
}

type FileServiceResponse struct {
	Data any
	Id   int
	Code string
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
		return FileServiceResponse{Code: "json parse error", Data: ""}
	}
	response := FileServiceResponse{}
	switch comobj.Con {
	case "delete":
		_, err := f.DeleteFile(comobj.Filepath, comobj.Filename)
		if err != nil {
			// return FileServiceResponse{code: "delete file error", data: ""}
			response.Code = "delete file error"
			response.Data = err.Error()
			return response
		}
		// return FileServiceResponse{code: "delete file success", data: comobj.filepath + "," + comobj.filename}
		response.Code = "delete file success"
		response.Data = comobj.Filepath + "," + comobj.Filename
	case "list":
		filelist := ListFilesAndDirs(comobj.Filepath)
		// return FileServiceResponse{code: "list file success", data: filelist}
		response.Code = "list file success"
		response.Data = filelist
		// 处理 list 命令
	case "download":
		// 处理 download 命令
		// return FileServiceResponse{code: "download file success", data: ""}
		response.Code = "download file success"
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
		response.Code = "unknown command"
		response.Data = ""
	}
	return response

}

func NewFileService(basePath string) *FileService {
	return &FileService{BasePath: basePath}
}

func ListFilesAndDirs(path string) []FileObject {
	if path == "" {
		drives := ListDiskDrives()
		filelist := []FileObject{}
		for _, drive := range drives {
			filelist = append(filelist, FileObject{Name: drive, IsDir: true, Path: drive})
		}
		return filelist
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil
	}
	filelist := []FileObject{}
	for _, entry := range entries {
		fileobj := FileObject{Name: entry.Name(), Path: path + "/" + entry.Name()}
		fileInfo, _ := entry.Info()
		if fileInfo != nil {
			fileobj.Size = fileInfo.Size()
			fileobj.ModTime = fileInfo.ModTime()
		}
		if entry.IsDir() {
			fileobj.IsDir = true
		} else {
			fileobj.IsDir = false
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
