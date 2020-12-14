package util

import (
	"io/ioutil"
	"os"
)

func PathExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func MkDir(path string, perm os.FileMode) {
	err := os.Mkdir(path, perm)
	Check(err, "cannot make dir: "+path)
}

func MkDirAll(path string, perm os.FileMode) {
	err := os.MkdirAll(path, perm)
	Check(err, "cannot make dir: "+path)
}

func MkFile(path string) *os.File {
	file, err := os.Create(path)
	Check(err, "cannot create file: "+path)
	return file
}

func WriteToFile(path string, content string) {
	file := MkFile(path)
	defer func() {
		_ = file.Close()
	}()
	_, err := file.Write([]byte(content))
	Check(err, "fail to write file: "+file.Name())
}

func ReadFileAll(path string) []byte {
	defer PanicRecover()
	file, err := os.Open(path)
	defer func() {
		_ = file.Close()
	}()
	Check(err, "cannot read file: "+path)
	content, readErr := ioutil.ReadAll(file)
	Check(readErr, readErr)
	return content
}

func OverWriteFile(path string, content []byte) {
	defer PanicRecover()
	file, openErr := os.OpenFile(path, os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	defer func() {
		_ = file.Close()
	}()
	Check(openErr, openErr)
	_, err := file.Write(content)
	Check(err, err)
}
