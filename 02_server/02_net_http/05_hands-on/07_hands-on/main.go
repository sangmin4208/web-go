package main

import (
	"bufio"
	"fmt"
	"net"
)

func main () {
	l, err := net.Listen("tcp",":8080")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept() 
		if err !=nil {
			fmt.Println(err)
			continue
		}
		go serve(conn)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
	}
}