package main

import (
	"fmt"
	"runtime"
	"net"
	"os"
)

func recvFile(conn net.Conn) {
	fmt.Println("new connection: ", conn.RemoteAddr())
	defer conn.Close()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var (
		port = "9090"
		remote = ":" + port
	)
	fmt.Println(remote)

	lis, err := net.Listen("tcp", remote)
    	defer lis.Close()

	if err != nil {
		fmt.Println("server listen failed!", err)
		os.Exit(-1)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("client connect failed!", err)
			continue
		}

		go recvFile(conn)
	}
}