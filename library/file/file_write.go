package file

import (
	"io"
	"os"
)

func (this *File) BackupFile(dstFile string) (res int64, err error) {
	src, err := os.Open(this.name)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
