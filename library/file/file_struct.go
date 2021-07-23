package file

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	. "github.com/rainmyy/easyDB/library/common"
	. "github.com/rainmyy/easyDB/library/strategy"
)

/**
*file struct
 */
type File struct {
	name     string
	filepath string
	size     int64
	content  []*TreeStruct
	mode     string
	isDir    bool
	/**
	* 文件数据类型有四种，ini型数据，yal型数据，json型数据，data型数据，默认数据类型为data型
	 */
	dataType int
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
	this.dataType = this.getDataType()
	return this
}

func FileInstance(name string) *File {
	fileObj := &File{}
	fileObj.getFilePath(name)
	fileObj = fileObj.fileInfo()
	if fileObj.isDir {
		return nil
	}
	return fileObj
}

func (this *File) getFilePath(fullname string) {
	fileAbs, err := filepath.Abs(fullname)
	if err != nil {
		return
	}
	fileName := path.Base(fullname)
	this.name = fileName
	this.fileAbs = fileAbs
}

func (this *File) getDataType() (dataType int) {
	fileSuffix := path.Ext(this.name)
	switch fileSuffix {
	case IniSuffix:
		dataType = IniType
	case JsonSuffix:
		dataType = JsonType
	case YamlSuffix:
		dataType = YamlType
	default:
		dataType = DataType
	}
	return
}

func (this *File) GetContent() []*TreeStruct {
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
