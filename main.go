package main

import (
	"crypto/rand"
	"flag"
	"github.com/cheggaaa/pb/v3"
	"io"
	"log"
	"os"
)

func main() {
	path := flag.String("o", "", "Output file")
	size := flag.Int64("s", 0, "Size file(Bytes)")
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
	err = process(file, *size)
	if err != nil {
		log.Fatal(err)
	}
}

func process(output *os.File, limit int64) error {
	reader := io.LimitReader(rand.Reader, limit)
	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(reader)
	_, err := io.Copy(output, barReader)
	if err != nil {
		return err
	}
	bar.Finish()
	return nil
}
