package ncbigrind

import (
	"github.com/indraniel/srasearch/sra"
	"github.com/indraniel/srasearch/utils"

	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"path"
)

func ProcessTarXMLs(
	tarfile string,
	accessionDB *map[string]*AccessionRecord,
	uploadsDB *map[string][]string,
) *map[string]*sra.SraItem {
	gzreader := utils.OpenGZFile(tarfile)
	defer gzreader.Close()

	tarReader := tar.NewReader(gzreader.Gzf)

	db := make(map[string]*sra.SraItem)

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
				sraItems := processXML(accessionDB, uploadsDB, name, buf)
				for _, si := range sraItems {
					db[si.Id] = si
				}
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

	return &db
}

func isXML(filename string) bool {
	if path.Ext(filename) == ".xml" {
		return true
	}
	return false
}

func processXML(
	accessionDB *map[string]*AccessionRecord,
	uploadsDB *map[string][]string,
	name string,
	contents *bytes.Buffer,
) []*sra.SraItem {
	sraItems := sra.NewSraItemsFromXML(name, contents.Bytes())
	for _, si := range sraItems {
		si.AddAttrFromAccessionRecords(accessionDB)
		si.AddAttrFromUploadRecords(uploadsDB)
	}

	return sraItems
}
