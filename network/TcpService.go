package network

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
)

type NetworkData struct {
	id     int
	ToAddr string
	ToPort int
	data   string
}

type TcpService struct {
	host    string
	port    int
	isInit  bool
	InChan  chan string
	OutChan chan string
}

func (this *TcpService) Init(addr string, port int) {
	this.InChan = make(chan string)
	this.OutChan = make(chan string)
}
func (this *TcpService) Run(addr string, port int) {
	this.host = addr
	this.port = port
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.host, this.port))
	if err != nil {
		fmt.Println("listen error:", err)
		panic(err)
	}
	defer listener.Close()
	var wg sync.WaitGroup
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			panic(err)

		}
		wg.Add(1)
		//conn传值
		go func() {
			defer wg.Done()
			buf := make([]byte, 1024)
			for {
				n, err := conn.Read(buf)
				if errors.Is(err, io.EOF) {
					break

				} else if err != nil {
					fmt.Println("read error:", err)
					panic(err)
				}
			}
			msg := string(buf[:n])
			fmt.Println("receive message:", msg)

		}()

	}
	wg.Wait()
}


func (this *TcpService) 