package main

import (
	"crypto/rand"
	"flag"
	"log"
	"os"
)

const bufferSize = 2048

func main() {
	path := flag.String("o", "", "Output file")
	size := flag.Int("s", 0, "Size file(Bytes)")
	flag.Parse()
	if *path == "" || *size == 0 {
		log.Fatal("Specify size(-s) and output(-o)")
	}
	log.Printf("Create file %s with %d bytes\n", *path, *size)
	file, err := os.OpenFile(*path, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	chunks := *size / bufferSize
	remain := *size - chunks*bufferSize

	buf := make([]byte, bufferSize)
	for i := 0; i < chunks; i++ {
		_, err = rand.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.Write(buf)
		if err != nil {
			log.Fatal(err)
		}
	}
	remainsBytes := make([]byte, remain)
	_, err = rand.Read(remainsBytes)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(remainsBytes)
	if err != nil {
		log.Fatal(err)
	}
}
