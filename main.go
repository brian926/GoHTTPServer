package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 8)
	for {
		bytesRead, err := file.Read(buffer)

		if bytesRead > 0 {
			fmt.Printf("read: %s\n", string(buffer[:bytesRead]))
		}

		if err != nil {
			if err == io.EOF {
				fmt.Println("End of the line...")
				break
			}
			log.Fatal(err)
		}
	}
}
