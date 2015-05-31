package makeuploads

import (
	"github.com/indraniel/srasearch/utils"

	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

func CreateRecentUploadsFile(uploads, outPath string, threshold int) {
	log.Println("Calculating number of entries")
	total := calculateTotalEntries(uploads)
	log.Printf("Got %d records\n", total)

	out := outputFileName(uploads, outPath, threshold)
	log.Println("Creating Output File:", out)

	start := total - threshold + 1 // +1 to account header line
	makeTSVFile(uploads, out, start)

	//	log.Println("Creating FOO -- threshold BAZ")
	log.Println("All Done!")
}

func makeTSVFile(uploads, outFile string, startLine int) {
	gzreader := utils.OpenGZFile(uploads)
	defer gzreader.Close()
	reader := bufio.NewReader(gzreader.Gzf)

	f, err := os.Create(outFile)
	if err != nil {
		log.Fatalf(
			"Trouble creating file '%s' : %s",
			outFile, err,
		)
	}
	defer f.Close()

	linenum := 0
	var e error
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			e = err
			break
		}
		linenum += 1

		if linenum == 1 || linenum >= startLine {
			_, err := f.WriteString(line)
			if err != nil {
				log.Fatalf(
					"[err] %s [%s] : %s",
					"Write error",
					f.Name(),
					err,
				)
			}
		}
	}

	if e != io.EOF {
		log.Fatalln("[err] reading line ", linenum, "in", gzreader.File.Name(), ":", e)
	}

	err = f.Sync()
	if err != nil {
		log.Fatalf(
			"%s [%s] : %s",
			"Trouble flushing output",
			f.Name(),
			err,
		)
	}
}

func calculateTotalEntries(uploads string) int {
	gzreader := utils.OpenGZFile(uploads)
	defer gzreader.Close()

	reader := bufio.NewReader(gzreader.Gzf)

	count := 0
	var e error
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			e = err
			break
		}
		count += 1
	}

	if e != io.EOF {
		log.Fatalln("[err] reading line ", count, "in", gzreader.File.Name(), ":", e)
	}

	return count
}

func outputFileName(uploads, outPath string, threshold int) string {
	ncbiFileName := path.Base(uploads)
	ncbiBaseName := strings.Split(ncbiFileName, ".")[0]

	ncbiBaseNameComponents := strings.Split(ncbiBaseName, "_")
	uploadDate := ncbiBaseNameComponents[len(ncbiBaseNameComponents)-1]

	outFile := fmt.Sprintf(
		"recent-%d-sra-uploads-%s.tsv",
		threshold,
		uploadDate,
	)

	fullPath := path.Join(outPath, outFile)
	return fullPath
}
