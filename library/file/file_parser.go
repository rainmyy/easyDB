package file

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"sync"
	"time"

	. "github.com/rainmyy/easyDB/library/common"
	. "github.com/rainmyy/easyDB/library/strategy"
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

// file size 1GB
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
	//if the file larger than 1GB,concurrently read and parse files
	if fileSize > defaultSize {
		return this.readFileByConcurrent(fi)
	} else {
		return this.readFileByGeneral(fi)
	}
}

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
	tree, err := parserDataFunc(this, b)
	if err != nil {
		return err
	}
	this.content = tree
	return nil
}

/**
* 并发读取,所有字符串按行分割， 暂不支持多行关联行数据
 */
func (this *File) readFileByConcurrent(fileObj *os.File) error {
	liensPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, 500*1024)
		return lines
	}}
	stringPool := sync.Pool{New: func() interface{} {
		lines := ""
		return lines
	}}
	slicePool := sync.Pool{New: func() interface{} {
		lines := make([]string, 100)
		return lines
	}}
	r := bufio.NewReader(fileObj)
	var wg sync.WaitGroup
	for {
		buf := liensPool.Get().([]byte)
		n, err := r.Read(buf)
		if n == 0 {
			if err != nil {
				fmt.Println(err)
				break
			}
			if err == io.EOF {
				break
			}
			return err
		}
		nextLine, err := r.ReadBytes('\n')
		if err != io.EOF {
			buf = append(buf, nextLine...)
		}
		wg.Add(1)
		go func() {
			ProcessChunk(buf, &liensPool, &stringPool, &slicePool)
			wg.Done()
		}()
	}

	wg.Wait()
	return nil
}

func ProcessChunk(chunk []byte, linesPool *sync.Pool, stringPool *sync.Pool, slicePool *sync.Pool) {
	var wg2 sync.WaitGroup
	logs := stringPool.Get().(string)
	logs = string(chunk)
	linesPool.Put(chunk)
	logsSlice := strings.Split(logs, "\n")
	stringPool.Put(logs)
	chunkSize := 100
	n := len(logsSlice)
	threadNo := n / chunkSize
	if n%chunkSize != 0 {
		threadNo++
	}
	length := len(logsSlice)
	for i := 0; i < length; i += chunkSize {
		wg2.Add(1)
		go func(s int, e int) {
			for i := s; i < e; i++ {
				text := logsSlice[i]
				if len(text) == 0 {
					continue
				}
				processLine(text)
			}
			wg2.Done()
		}(i*chunkSize, int(math.Min(float64((i+1)*chunkSize), float64(len(logsSlice)))))
	}
}

func processLine(text string) {
	logParts := strings.SplitN(text, ",", 2)
	logCreationTimeString := logParts[0]
	_, err := time.Parse("2006-01-  02T15:04:05.0000Z", logCreationTimeString)
	if err != nil {
		fmt.Println("\n Could not able to parse the time:%s for log: %v", logCreationTimeString, text)
		return
	}
	//if logCreateTime.After(start) && logCreateTime.Before(end) {
	//	fmt.Println(text)
	//}
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
