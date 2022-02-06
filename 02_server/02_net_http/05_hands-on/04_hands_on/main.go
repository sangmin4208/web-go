package main

import (
	"fmt"
	"io"
	"net"
)

func main(){
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}	
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handle(conn)
	}
}

func handle (conn net.Conn){
	io.WriteString(conn, "I see you connected")
	conn.Close()
}