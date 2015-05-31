package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
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

func FindFile(pattern string) (string, error) {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return "", err
	}

	if len(matches) == 0 {
		err := fmt.Errorf("Found no files with pattern : %s", pattern)
		return "", err
	}

	if len(matches) > 1 {
		err := fmt.Errorf("Found multiple files with pattern : %s", pattern)
		return "", err
	}

	return matches[0], nil
}

func FileModificationTime(file string) (time.Time, error) {
	finfo, err := os.Stat(file)
	if err != nil {
		e := fmt.Errorf("[%s] File Stat error : %s", file, err)
		return time.Unix(0, 0), e
	}

	mtime := finfo.ModTime()
	return mtime, nil
}
