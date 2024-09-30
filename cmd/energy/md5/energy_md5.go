package main

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		println("args len != 2")
		return
	}
	wd, _ := os.Getwd()
	file := os.Args[1]
	filePath := filepath.Join(wd, file)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		println(err.Error())
		return
	}
	hash := md5.New()
	hash.Write(data)
	md5String := hex.EncodeToString(hash.Sum(nil))
	md5File, err := os.OpenFile("md5.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		println(err.Error())
		return
	}
	defer md5File.Close()
	md5File.WriteString(md5String + "  " + file + "\r\n")
	println("file:", file, "md5:", md5String)
}
