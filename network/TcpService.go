package network

import (
	"fmt"
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

func (t *TcpService) Init(addr string, port int) {
	t.InChan = make(chan string)
	t.OutChan = make(chan string)
}
func (t *TcpService) Run(addr string, port int) {
	t.host = addr
	t.port = port
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", t.host, t.port))
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
			break

		}
		wg.Add(1)
		//conn传值
		go t.StartReading(conn)

	}
	wg.Wait()
}

func (t *TcpService) StartReading(conn net.Conn) {
	buf := make([]byte, 1024)
	data := make([]byte, 1024)
	for {
		n, _ := conn.Read(buf)
		if n > 0 {
			data = append(data, buf[:n]...)
		} else if len(data) > 0 {
			//数据处理
		}

	}
}
