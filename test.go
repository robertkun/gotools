package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
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
