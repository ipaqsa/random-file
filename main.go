package main

import (
	"crypto/rand"
	"flag"
	"github.com/cheggaaa/pb/v3"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	path := flag.String("o", "", "Output file")
	size := flag.String("s", "", "Size file(B[bytes], K[Kb], M[Mb], G[Gb])")
	flag.Parse()
	if *path == "" || *size == "" {
		log.Fatal("Specify size(-s) and output(-o)")
	}
	log.Printf("Create file %s with %s\n", *path, *size)
	file, err := os.OpenFile(*path, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = process(file, parseSize(*size))
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

func parseSize(size string) int64 {
	last := size[len(size)-1]
	rawSize, err := strconv.ParseInt(size[:len(size)-1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	var mult int64 = 0
	switch last {
	case 'B':
		mult = 1
	case 'K':
		mult = 1024
	case 'M':
		mult = 1024 * 1024
	case 'G':
		mult = 1024 * 1024 * 1024
	}
	return rawSize * mult
}
