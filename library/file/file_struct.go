package file

import (
	"os"
	"time"
)

/**
*file struct
 */
type File struct {
	name    string
	size    int64
	content []byte
	mode    string
	isDir   bool
	modTime time.Time
}

func (this *File) fileInfo() *File {
	fi, err := os.Stat(this.name)
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

func FileInstance(name string) *File {
	fileObj := &File{name: name}
	fileObj = fileObj.fileInfo()
	return fileObj
}

