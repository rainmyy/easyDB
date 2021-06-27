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
	fpStr := fmt.Sprintf("%s/%s", this.filepath, this.name)
	fileAbs, err := filepath.Abs(fpStr)
	if err != nil {
		return
	}
	this.fileAbs = fileAbs
}
