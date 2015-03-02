package sradump

import (
	"archive/tar"
	"github.com/indraniel/srasearch/sra"
	"github.com/indraniel/srasearch/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

func ProcessTarXMLs(
	tarfile string,
	db *map[string]*sra.AccessionRecord,
	outFile string,
) {
	f, gzf := utils.OpenGZFile(tarfile)
	defer utils.CloseGZFile(f, gzf)

	tarReader := tar.NewReader(gzf)

	outPtr, err := os.Create(outFile)
	if err != nil {
		log.Fatal("Trouble opening %s for writing : %s\n", outFile, err)
	}
	defer outPtr.Close()

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
			//			fmt.Println("(", i, ")", "Name:", name)
			if isXML(name) {
				buf := new(bytes.Buffer)
				io.Copy(buf, tarReader)
				processXML(db, outPtr, name, buf)
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
		//		if i == 100 {
		//			os.Exit(0)
		//		}
	}
}

func isXML(filename string) bool {
	if path.Ext(filename) == ".xml" {
		return true
	}
	return false
}

func processXML(
	db *map[string]*sra.AccessionRecord,
	outPtr *os.File,
	name string,
	contents *bytes.Buffer,
) {
	sraItems := sra.NewSraItemsFromXML(name, contents.Bytes())
	for _, si := range sraItems {
		si.AddAttrFromAccessionRecords(db)
		json, err := json.Marshal(si)
		if err != nil {
			log.Fatal("Trouble encoding '%s' into json: %s\n",
				name, err)
		}

		line := strings.Join([]string{si.Id, string(json)}, ",")

		_, err = outPtr.WriteString(line)
		checkWrite(outPtr, err)
		_, err = outPtr.Write([]byte("\n"))
		checkWrite(outPtr, err)
	}
}

func checkWrite(f *os.File, e error) {
	if e != nil {
		log.Fatal("Trouble writing to '%s' : %s\n", f.Name(), e)
	}
}
