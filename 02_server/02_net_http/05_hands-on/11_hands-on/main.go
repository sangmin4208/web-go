package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main(){
	li,err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn,err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go serve(conn)
	}
}

func serve(c net.Conn){
	defer c.Close()
	var i int
	var rMethod, rURI string
	scanner := bufio.NewScanner(c) 
	for scanner.Scan(){
		ln := scanner.Text()
		if i == 0 {
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD:", rMethod)
			fmt.Println("URI:", rURI)
		}
		if ln == "" {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		i++
	}
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Code Gangsta</title>
		</head>
		<body>
			<h1>"HOLY COW THIS IS LOW LEVEL"</h1>
		</body>
		</html>
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}