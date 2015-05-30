package utils

import (
	"bufio"
	"compress/gzip"
	"io/ioutil"
	"log"
	"os"
)

type GZReader struct {
	File *os.File
	Gzf  *gzip.Reader
}

func OpenGZFile(filename string) GZReader {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	gzf, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}

	gzreader := GZReader{File: f, Gzf: gzf}
	return gzreader
}

func (g GZReader) Close() {
	g.File.Close()
	g.Gzf.Close()
}

type GZWriter struct {
	f   *os.File
	gf  *gzip.Writer
	buf *bufio.Writer
}

func NewGZWriter(filename string) (GZWriter, error) {
	f, err := os.Create(filename)
	if err != nil {
		return GZWriter{}, nil
	}

	gf := gzip.NewWriter(f)
	buf := bufio.NewWriter(gf)

	gzw := GZWriter{f, gf, buf}
	return gzw, nil
}

func (gz GZWriter) Write(data []byte) error {
	_, err := gz.buf.Write(data)
	return err
}

func (gz GZWriter) Close() {
	gz.buf.Flush()
	gz.gf.Close()
	gz.f.Close()
}

func CompressFile(src string, dst string) error {
	gz, err := NewGZWriter(dst)
	if err != nil {
		return err
	}
	defer gz.Close()

	log.Println("Writing compress data to:", gz.f.Name())
	in, err := os.Open(src)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	err = gz.Write(data)
	if err != nil {
		return err
	}

	return nil
}
