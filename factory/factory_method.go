package factory

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// 当对象的创建逻辑比较复杂，不只是简单的 new 一下就可以
// 而是要组合其他类对象，做各种初始化操作的时候
// 推荐使用工厂方法模式，将复杂的创建逻辑拆分到多个工厂类中，让每个工厂类都不至于过于复杂
// 示例：1. 选择解析对象类型 map slice
//       2. 再选择解析类型 json yaml
func NewIConfigParserMap(typ string) IConfigParserMap {
	var configParserMap IConfigParserMap
	switch typ {
	case "json":
		configParserMap = &JsonConfigParse{}
	case "yaml":
		configParserMap = &YamlConfigParse{}
	default:
		configParserMap = nil
	}
	return configParserMap
}

func NewIConfigParserSlice(typ string) IConfigParserSlice {
	var configParserMap IConfigParserSlice
	switch typ {
	case "json":
		configParserMap = &JsonConfigParse{}
	case "yaml":
		configParserMap = &YamlConfigParse{}
	default:
		configParserMap = nil
	}
	return configParserMap
}

type IConfigParserMap interface {
	ParseMap([]byte) (map[string]interface{}, error)
}

type IConfigParserSlice interface {
	ParseSlice([]byte) ([]interface{}, error)
}

type JsonConfigParse struct {
}

func (j *JsonConfigParse) ParseMap(data []byte) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	err := json.Unmarshal(data, &obj)
	return obj, err
}

func (j *JsonConfigParse) ParseSlice(data []byte) ([]interface{}, error) {
	obj := make([]interface{}, 0)
	err := json.Unmarshal(data, &obj)
	return obj, err
}

type YamlConfigParse struct {
}

func (y *YamlConfigParse) ParseMap(data []byte) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	err := yaml.Unmarshal(data, &obj)
	return obj, err
}

func (y *YamlConfigParse) ParseSlice(data []byte) ([]interface{}, error) {
	obj := make([]interface{}, 0)
	err := yaml.Unmarshal(data, &obj)
	return obj, err
}
