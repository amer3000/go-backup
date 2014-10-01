package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
	"flag"
)

//https://github.com/mrjbq7/re-factor/blob/master/dupe/dupe.go

var rootDir string = "."

func main() {

        //take CLI input
	dirPtr := flag.String("path","."," string")
	svrPtr := flag.String("server","www.softlayer.com"," string")
	frqPtr := flag.Int("bkp-interval", 60, "backup interval in hours")
	comprsPtr := flag.Bool("compression", false, "a bool")
	encrptPtr := flag.Bool("encryption", false, "a bool")
	flag.Parse()
	dir_path := *dirPtr
	server := *svrPtr
	frequency := *frqPtr
	compress := *comprsPtr
	encrypt := *encrptPtr
	//I am not including bakcup run time since 
	//having both frequency and backup run time 
	///does not make sense


	fmt.Println("Hello user your inputs are  ",dir_path,server, frequency,compress,encrypt)
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
