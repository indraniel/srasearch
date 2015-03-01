package utils

import (
	"compress/gzip"
	"log"
	"os"
)

func OpenGZFile(filename string) (*os.File, *gzip.Reader) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	gzf, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}

	return f, gzf
}

func CloseGZFile(f *os.File, gzf *gzip.Reader) {
	f.Close()
	gzf.Close()
}
