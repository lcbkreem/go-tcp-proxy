package main

import (
	"io"
	"log"
	"net"
)

func main() {
	//1. listen for proxy requests
	listener, err := net.Listen("tcp", ":1111")
	if err != nil {
		log.Fatalln("Unable to bind on port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		//2. handle that request
		go handleConn(conn)
	}
}

func handleConn(src net.Conn) {
	// 1. Connect to the target server
	dst, err := net.Dial("tcp", "www.google.com:443")
	if err != nil {
		log.Fatalln("Unable to connect to target server")
	}

	defer dst.Close()

	// 2. Copy the src output to the destination connection
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	// 3. Copy the response from target server connection to src
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}
