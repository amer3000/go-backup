package main

import (
	"crypto/md5"
	"encoding/hex"
	// "flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var rootDir string = "./testData"
var backupDir string = "/cygdrive/z/testBackup"

type FileHash struct {
	fullpath       string
	hash           string
	lastModifyTime time.Time
}

var fileHashes []FileHash

func main() {

	//take CLI input
/*	dirPtr := flag.String("path", ".", " string")
	svrPtr := flag.String("server", "www.softlayer.com", " string")
	frqPtr := flag.Int("bkp-interval", 60, "backup interval in hours")
	comprsPtr := flag.Bool("compression", false, "a bool")
	encrptPtr := flag.Bool("encryption", false, "a bool")
	flag.Parse()
	dir_path := *dirPtr
	server := *svrPtr
	frequency := *frqPtr
	compress := *comprsPtr
	encrypt := *encrptPtr
*/	//I am not including bakcup run time since
	//having both frequency and backup run time
	///does not make sense

//	fmt.Println("Hello user your inputs are  ", dir_path, server, frequency, compress, encrypt)

	t0 := time.Now()
	fmt.Printf("\nIt begins at %v", t0)
	
	/*FO, err := os.Create("backup.txt")
	if err != nil {
		panic(err)
	}
	defer FO.Close()
*/
	filepath.Walk(rootDir, VisitFile)

/*	for _, fh := range fileHashes {
		FO.WriteString(fh.fullpath + ", " + fh.hash + ", " + fh.lastModifyTime.String() + "\n")
	}
*/	t1 := time.Now()

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
		fullpath, _ = filepath.Abs(fullpath)
		hash := MD5OfFile(fullpath)
		fileHashes = append(fileHashes, FileHash{fullpath, hex.EncodeToString(hash), f.ModTime()})
	}
	return nil
}

func FileCopy(filename string) {
	
}