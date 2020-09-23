package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Print("Error in Listen tcp connection. ", err.Error())
		return
	}
	for true {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print("Error in accepting. ", err.Error())
			return
		}
		go dealwithConn(conn)
	}
}

func dealwithConn(conn net.Conn) {
	for true {
		bytes := make([]byte, 512)
		len, err := conn.Read(bytes)
		if err != nil {
			fmt.Print("Error in Read. ", err.Error())
			return
		}
		fmt.Printf("Receive data :%v \n", string(bytes[:len]))
	}
}
