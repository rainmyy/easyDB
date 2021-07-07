package file

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

/**
*file struct
 */
type File struct {
	name     string
	filepath string
	size     int64
	content  []byte
	mode     string
	isDir    bool
	modTime  time.Time
	fileAbs  string
}

func (this *File) fileInfo() *File {
	fi, err := os.Stat(this.fileAbs)
	if err != nil {
		return this
	} else if os.IsNotExist(err) {
		return this
	}
	this.size = fi.Size()
	this.mode = fi.Mode().String()
	this.modTime = fi.ModTime()
	this.isDir = fi.IsDir()
	return this
}

func FileInstance(name string, filepath string) *File {
	fileObj := &File{name: name, filepath: filepath}
	fileObj.getFilePath()
	fileObj = fileObj.fileInfo()
	if fileObj.isDir {
		return nil
	}
	fileObj.readFile()
	return fileObj
}

func (this *File) getFilePath() {
	fileAbs, err := filepath.Abs(filepath.Join(this.filepath, this.name))
	if err != nil {
		return
	}
	this.fileAbs = fileAbs
}

func (this *File) GetContent() []byte {
	return this.content
}

func (this *File) createDir(fileName string) bool {
	if this.checkFileExist(fileName) {
		return true
	}
	paths, _ := filepath.Split(fileName)
	if this.checkFileExist(paths) {
		return true
	}
	err := os.MkdirAll(paths, os.ModePerm)
	if err != nil {
		fmt.Println("create file: " + fileName + " fail, msg:" + fmt.Sprintf("%s", err))
		return false
	}
	return true
}

func (this *File) checkFileExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
