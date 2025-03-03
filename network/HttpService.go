package network

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../service"
)

type HttpService struct {
	Server *http.Server
}

func HttpServiceRun(port int) *HttpService {
	h := new(HttpService)
	h.Server = new(http.Server)

	h.SetRoute(h.Server)
	fmt.Println("0.0.0.0:" + strconv.Itoa(port))
	h.Server.Addr = "0.0.0.0:" + strconv.Itoa(port)

	go h.Server.ListenAndServe()

	return h
}

func (h *HttpService) SetRoute(s *http.Server) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))

	})
	mux.HandleFunc("/file/list", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()

		// 获取参数值
		path := params.Get("path") // 获取单个参数值
		if path == "" {
			path = "D:/" // 默认路径
		}

		// 调用 ListFilesAndDirs 获取文件列表
		filelist := service.ListFilesAndDirs(path)

		// 将文件列表转换为 JSON 字符串
		filelistStr, _ := json.Marshal(filelist)

		// 设置响应头
		w.Header().Set("Content-Type", "application/json")

		// 写入响应
		w.Write([]byte(filelistStr))

	})

	// s.Handler = mux
	s.Handler = h.corsMiddleware(mux)

}
func (h *HttpService) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置 CORS 头
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 处理预检请求（OPTIONS）
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
