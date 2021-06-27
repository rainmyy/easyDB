/***
*read file
 */
package file

import (
	"bufio"
	"io"
	"os"
)

//file size 1GB
var defaultSize = 1 << 30

func (this *File) readFile() {
	fileName := this.fileAbs
	fi, err := os.Open(fileName)
	defer fi.Close()
	if err != nil {
		return
	}
	fileSize := this.size
	if fileSize == 0 {
		fiStat, err := fi.Stat()
		if err != nil {
			return
		}
		fileSize = fiStat.Size()
	}
	//大于1GB的文件并行读取
	if fileSize > int64(defaultSize) {
		this.readFileByConcurrent(fi)
	} else {
		this.readFileByGeneral(fi)
	}
}

//
func (this *File) readFileByGeneral(fileObj *os.File) {
	if fileObj == nil {
		return
	}
	r := bufio.NewReader(fileObj)
	b := make([]byte, this.size)
	for {
		_, err := r.Read(b)
		if err != nil && err == io.EOF {
			panic(err)
			break
		}
	}
	if len(b) > 0 {
		this.content = b
	}
}

func (this *File) readFileByConcurrent(fileObj *os.File) {

}
