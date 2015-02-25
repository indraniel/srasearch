package jdoc

import (
	"archive/tar"
	"github.com/indraniel/srasearch/sra"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func ProcessNCBITarFile(tarfile string) {
	f, err := os.Open(tarfile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer gzf.Close()

	tarReader := tar.NewReader(gzf)

	i := 0
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		name := header.Name

		switch header.Typeflag {
		case tar.TypeDir:
			continue
		case tar.TypeReg:
			fmt.Println("(", i, ")", "Name:", name)
			if isXML(name) {
				buf := new(bytes.Buffer)
				io.Copy(buf, tarReader)
				processXML(name, buf)
			}
		default:
			fmt.Printf("%s: %c %s %s\n",
				"Yikes! Unable to figure out type",
				header.Typeflag,
				"in file",
				name,
			)
		}
		i++
	}
}

func isXML(filename string) bool {
	if path.Ext(filename) == ".xml" {
		return true
	}
	return false
}

func processXML(name string, contents *bytes.Buffer) {
	si := sra.NewSraItemFromXML(name, contents.Bytes())
	fmt.Println("---")
	io.Copy(os.Stdout, bytes.NewBufferString(si.Data.String()))
	fmt.Println("---")
	os.Exit(0)
}
