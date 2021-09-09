package factory

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

type IConfigParser interface {
	Parse([]byte) (interface{}, error)
}

type JsonConfigParseInterface struct {
}

func (j *JsonConfigParseInterface) Parse(data []byte) (interface{}, error) {
	obj := make(map[string]interface{})
	err := json.Unmarshal(data, &obj)
	return obj, err
}

type YamlConfigParseInterface struct {
}

func (y *YamlConfigParseInterface) Parse(data []byte) (interface{}, error) {
	obj := make(map[string]interface{})
	err := yaml.Unmarshal(data, &obj)
	return obj, err
}

// 根据参数去选择类型
// 适用于简单的对象构建
// 缺点是新增类型需要改动原有代码
func NewIConfigParser(typ string) IConfigParser {
	var configParser IConfigParser
	switch typ {
	case "json":
		configParser = &JsonConfigParseInterface{}
	case "yaml":
		configParser = &YamlConfigParseInterface{}
	default:
		configParser = nil
	}
	return configParser
}
