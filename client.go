package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"io"
)

func SendFile(conn net.Conn) {
	fileName := "C:\\Robert\\日志分析\\tools_go\\vdn_sqlInterface\\a.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("file size=", size)

	// define read block size = 2
	buf := make([]byte, 2)
	for {
		length, err := file.Read(buf)
		if err != nil {
		    if err == io.EOF {
			break
		    } else {
			fmt.Println("Read file error!", err)
			return
		    }
		}

		fmt.Println(length, string(buf))
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var (
		host   = "127.0.0.1"
		port   = "9090"
		remote = host + ":" + port
	)

	fmt.Println(remote)
	conn, err := net.Dial("tcp", remote)
	if err != nil {
		fmt.Println("connect server failed!.")
		os.Exit(-1)
		return
	}
	defer conn.Close()

	fmt.Println(0, "connect ok! sending file...")
	SendFile(conn)
}
