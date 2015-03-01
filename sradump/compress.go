package sradump

import (
	"bufio"
	"compress/gzip"
	"io/ioutil"
	"log"
	"os"
)

type GZipFile struct {
	f   *os.File
	gf  *gzip.Writer
	buf *bufio.Writer
}

func CreateGZ(filename string) (GZipFile, error) {
	f, err := os.Create(filename)
	if err != nil {
		return GZipFile{}, nil
	}

	gf := gzip.NewWriter(f)
	buf := bufio.NewWriter(gf)

	gzf := GZipFile{f, gf, buf}
	return gzf, nil
}

func (gz GZipFile) WriteGZ(data []byte) error {
	_, err := gz.buf.Write(data)
	return err
}

func (gz GZipFile) CloseGZ() {
	gz.buf.Flush()
	gz.gf.Close()
	gz.f.Close()
}

func CompressDumpFile(src string, dst string) error {
	gz, err := CreateGZ(dst)
	if err != nil {
		return err
	}

	log.Println("Writing compress data to:", gz.f.Name())
	in, err := os.Open(src)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	err = gz.WriteGZ(data)
	if err != nil {
		return err
	}

	return nil
}

func openGZFile(filename string) (*os.File, *gzip.Reader) {
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

func closeGZFile(f *os.File, gzf *gzip.Reader) {
	f.Close()
	gzf.Close()
}
