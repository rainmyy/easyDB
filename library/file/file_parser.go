package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
* 解析数据，将数据解析成树形结构进行存储
 */
func (this *File) Parser(objType int) error {
	err := this.readFile()
	if err != nil {
		return err
	}
	return nil
}

//file size 1GB
var defaultSize int64 = 1 << 30

func (this *File) readFile() error {
	fileName := this.fileAbs
	fi, err := os.Open(fileName)
	defer fi.Close()
	if err != nil {
		return err
	}
	fileSize := this.size
	if fileSize == 0 {
		fiStat, err := fi.Stat()
		if err != nil {
			return err
		}
		fileSize = fiStat.Size()
	}
	//大于1GB的文件并行读取
	if fileSize > defaultSize {
		return this.readFileByConcurrent(fi)
	} else {
		return this.readFileByGeneral(fi)
	}
}

//
func (this *File) readFileByGeneral(fileObj *os.File) error {
	if fileObj == nil {
		return fmt.Errorf("file is nil")
	}
	r := bufio.NewReader(fileObj)
	b := make([]byte, this.size)
	for {
		_, err := r.Read(b)
		if err != nil && err == io.EOF {
			break
		}
	}
	dataType := this.dataType
	tree, err := parserDataFunc(this, dataType, b)
	if err != nil {
		return err
	}
	this.content = tree
	return nil
}

/**
* 并行读取
 */
func (this *File) readFileByConcurrent(fileObj *os.File) error {
	return nil
}
