package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)
	go func() {
		defer f.Close()
		defer close(out)
		line := ""
		for {
			buffer := make([]byte, 8)
			bytesRead, err := f.Read(buffer)
			if err != nil {
				break
			}

			buffer = buffer[:bytesRead]

			if i := bytes.IndexByte(buffer, '\n'); i != -1 {
				line += string(buffer[:i])
				out <- line
				buffer = buffer[i+1:]
				line = ""
			}

			line += string(buffer)
		}
		if len(line) != 0 {
			out <- line
		}
	}()
	return out
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error", "error", err)
		}

		linesChan := getLinesChannel(conn)

		for line := range linesChan {
			fmt.Println(line)
		}
		fmt.Println("Connection to ", conn.RemoteAddr(), "closed")
	}
}
