package main

import (
	"bufio"
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

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	defer conn.Close()
	io.WriteString(conn, "I see you connected")
}