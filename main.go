package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}

	line := ""
	for {
		buffer := make([]byte, 8)
		bytesRead, err := file.Read(buffer)

		buffer = buffer[:bytesRead]

		if i := bytes.IndexByte(buffer, '\n'); i != -1 {
			line += string(buffer[:i])
			//fmt.Printf("read: %s\n", line)
			buffer = buffer[i+1:]
			line = ""
		}

		line += string(buffer)

		if err != nil {
			if err == io.EOF {
				fmt.Println("End of the line...")
				break
			}
			log.Fatal(err)
		}
	}
}
