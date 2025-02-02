package network

import "net"

type ConnList struct {
	conns []net.Conn
}

func (cl *ConnList) Init() {
	cl.conns = make([]net.Conn, 10)
}

func (cl *ConnList) Add(conn net.Conn) {
	cl.conns = append(cl.conns, conn)
}

func (cl *ConnList) Del(conn net.Conn) {
	for i, v := range cl.conns {
		if v == conn {
			cl.conns = append(cl.conns[:i], cl.conns[i+1:]...)
		}
	}
}

func (cl *ConnList) GetAll() []net.Conn {
	return cl.conns
}

func (cl *ConnList) Get(index int) net.Conn {
	return cl.conns[index]
}

func (cl *ConnList) Len() int {
	return len(cl.conns)
}

func (cl *ConnList) Clear() {
	cl.conns = make([]net.Conn, 0)
}

func (cl *ConnList) GetIndex(conn net.Conn) int {
	for i, v := range cl.conns {
		if v == conn {
			return i
		}
	}
	return -1
}
