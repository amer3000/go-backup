package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

//https://github.com/mrjbq7/re-factor/blob/master/dupe/dupe.go

var rootDir string = "."

func main() {
	fmt.Println("Hello World!!!!!")
	t0 := time.Now()
	filepath.Walk(rootDir, VisitFile)
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}

func MD5OfFile(fullpath string) []byte {
	if contents, err := ioutil.ReadFile(fullpath); err == nil {
		md5sum := md5.New()
		md5sum.Write(contents)
		return md5sum.Sum(nil)
	}
	return nil
}

func VisitFile(fullpath string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !f.IsDir() {
		hash := MD5OfFile(fullpath)
		fmt.Printf("%s    %x\n", fullpath, hash)
	}
	return nil
}
