package file

import "github.com/easydb/library/common"

/**
* 解析
 */
func (this *File) Parser(obj interface{}) {
	common.ParserIniContent(this.content)
}
