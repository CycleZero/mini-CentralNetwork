package network

import (
	"fmt"
	"net/http"
	"strconv"
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
		//TODO:filelist

	})

	s.Handler = mux

}
