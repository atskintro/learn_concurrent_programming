package main

import (
	"io"
	"log"
	"os"
	"time"
)

func readContent(reader io.Reader) {
	content := ""
	buffer := make([]byte, 1024)
	for {
		bytes_read, err := reader.Read(buffer)
		log.Printf("Bytes Read: %d\n", bytes_read)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error reading the file: %v", err)
		}
		content += string(buffer)
	}
	log.Print(content)
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("filename should be given")
	}

	for _, name := range os.Args[1:] {
		file, err := os.Open(name)
		if err != nil {
			log.Printf("Error in reading file:%s err:%v", name, err)
		}
		go readContent(file)
	}
	time.Sleep(2 * time.Second)
}
