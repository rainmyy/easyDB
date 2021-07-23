package file

import (
	"log"

	. "github.com/rainmyy/easyDB/library/strategy"
	"gopkg.in/yaml.v2"
)

func ParserYamlContent(data []byte) ([]*TreeStruct, error) {
	yamlContent := `
field1: 小韩说课
field2:
  field3: value
  field4: [42, 1024]
`
	// 存储解析数据
	result := make(map[string]interface{})
	// 执行解析
	err := yaml.Unmarshal([]byte(yamlContent), &result)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return nil, nil
}
