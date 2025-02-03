package network

import (
	"fmt"
	"net"
	"sync"

	constformat "../ConstFormat"
)

type TcpService struct {
	host   string
	port   int
	isInit bool

	OutChan chan constformat.NetCommandPackage
}

func (t *TcpService) Init(addr string, port int) {
	t.host = addr
	t.port = port
	t.OutChan = make(chan constformat.NetCommandPackage, 100)
	t.isInit = true
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
		go t.hundleConnect(conn)

	}
	wg.Wait()
}

func (t *TcpService) Reader(conn net.Conn, result chan string) {
	buf := make([]byte, 1024)
	data := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if n > 0 {
			data = append(data, buf[:n]...)
		} else if len(data) > 0 {
			result <- string(data)
		} else if err != nil {
			fmt.Println("read error:", err)
			break
		}

	}
}

func (t *TcpService) Sender(conn net.Conn, datach chan string) {
	for {
		data := <-datach
		n, err := conn.Write([]byte(data))
		if n != len(data) {
			fmt.Println("send error: n < length")
		}
		if err != nil {
			fmt.Println("send error:", err)
			break
		}
	}
}

func (t *TcpService) hundleConnect(conn net.Conn) {
	defer conn.Close()
	readResult := make(chan string, 10)
	sendchannel := make(chan string, 10)
	go t.Reader(conn, readResult)
	go t.Sender(conn, sendchannel)
	// for {
	// 	select {
	// 	case data := <-readResult:
	// 		fmt.Println("read:", data)

	// 		t.OutChan <- data
	// 	case data := <-t.OutChan:
	// 		fmt.Println("send:", data)
	// 		sendchannel <- data
	// 	}
	// }

}
