package file

import (
	"bufio"
	"fmt"
	"io"
	"os"

	. "github.com/rainmyy/easyDB/library/common"
	. "github.com/rainmyy/easyDB/library/strategy"
)

/**
* 解析数据，将数据解析成树形结构进行存储
 */
func (f *File) Parser(objType int) error {
	err := f.readFile()
	if err != nil {
		return err
	}
	return nil
}

// file size 1GB
var defaultSize int64 = 1 << 30

func (f *File) readFile() error {
	fileName := f.fileAbs
	fi, err := os.Open(fileName)
	defer fi.Close()
	if err != nil {
		return err
	}
	fileSize := f.size
	if fileSize == 0 {
		fiStat, err := fi.Stat()
		if err != nil {
			return err
		}
		fileSize = fiStat.Size()
	}
	//if the file larger than 1GB,concurrently read and parse files
	if fileSize > defaultSize {
		return f.readFileByConcurrent(fi)
	} else {
		return f.readFileByGeneral(fi)
	}
}

func (f *File) readFileByGeneral(fileObj *os.File) error {
	if fileObj == nil {
		return fmt.Errorf("file is nil")
	}
	r := bufio.NewReader(fileObj)
	b := make([]byte, f.size)
	for {
		_, err := r.Read(b)
		if err != nil && err == io.EOF {
			break
		}
	}
	tree, err := parserDataFunc(f, b)
	if err != nil {
		return err
	}
	f.content = tree
	return nil
}

/**
* 并发读取,所有字符串按行分割， 暂不支持多行关联行数据
 */
func (f *File) readFileByConcurrent(fileObj *os.File) error {
	return nil
}

/**
* 所有的数
 */
func parserDataFunc(file *File, data []byte) ([]*TreeStruct, error) {
	var objType = file.GetDataType()

	switch objType {
	case IniType:
		return ParserIniContent(data)
	case YamlType:
		return ParserYamlContent(data)
	case JsonType:
		return ParserJsonContent(data)
	default:
		return ParserContent(data)
	}
}
