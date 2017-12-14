package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

type File struct {

}

func NewFile() *File {
	return &File{}
}

// get file contents
func (f *File) GetFileContents(filePath string) (content string, err error) {
	defer func(err *error) {
		e := recover()
		if e != nil {
			*err = fmt.Errorf("%s", e)
		}
	}(&err)
	bytes, err := ioutil.ReadFile(filePath)
	content = string(bytes)
	return
}

// file or path is exists
func (f *File) PathIsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}