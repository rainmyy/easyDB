/***
*read file
 */
package file

import (
	"os"
)

func (this *File) readFile() {
	fileName := this.name
	fi, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer fi.Close()

	fileSize := this.size
}
