package utils

import (
	"io/ioutil"
	"log"
	"os"
)

func CheckFileExists(file string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.Fatalf(
			"Could not find '%s' on file system: %s",
			file, err,
		)
	}
}

func MakeTmpFile() (tmpdir, tmpfile string) {
	tmpdir, err := ioutil.TempDir(os.TempDir(), "sra-dump")
	if err != nil {
		log.Fatal("Trouble making temp dir:", err)
	}

	f, err := ioutil.TempFile(tmpdir, "sra-tmp-dump")
	if err != nil {
		log.Fatal("Trouble making temp file:", err)
	}
	defer f.Close()

	tmpfile = f.Name()
	return
}

func CheckWrite(f *os.File, e error) {
	if e != nil {
		log.Fatal("Trouble writing to '%s' : %s\n", f.Name(), e)
	}
}
