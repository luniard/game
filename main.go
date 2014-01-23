package main

import (
	"fmt"
	"io"
	"net"

	gamenet "github.com/luniard/game/net"
)

const RECV_BUF_LEN = 1024

var (
	Pipeline *gamenet.Pipe
)

func init() {
	Pipeline = gamenet.NewPipe()
}

func main() {
	listener, err := net.Listen("tcp", ":8080") //侦听在6666端口
	if err != nil {
		panic("error listening:" + err.Error())
	}
	fmt.Println("Starting the server")

	for {
		conn, err := listener.Accept() //接受连接
		if err != nil {
			panic("Error accept:" + err.Error())
		}
		fmt.Println("Accepted the Connection :", conn.RemoteAddr())
		go EchoServer(conn)
	}
}

func EchoServer(conn net.Conn) {
	buf := make([]byte, RECV_BUF_LEN)
	defer conn.Close()

	for {
		_, err := conn.Read(buf)
		switch err {
		case nil:
			data := Pipeline.Handle(buf)
			// conn.Write(buf[0:n])
			// fmt.Println("send to client", data)
			conn.Write(data)
		case io.EOF:
			fmt.Printf("Warning: End of data: %s \n", err)
			return
		default:
			fmt.Printf("Error: Reading data : %s \n", err)
			return
		}
	}
}
