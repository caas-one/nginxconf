package nginx

import (
	"fmt"
	"io/ioutil"

	"github.com/json-iterator/go"
)

// 元Block，最内层的Block
type MetaBlock struct {
	Line      int      `json:"line"`
	Args      []string `json:"args"`
	Directive string   `json:"directive"`
}

// 倒数第二层的Block
type InmostBlock struct {
	Line       int         `json:"line"`
	Args       []string    `json:"args"`
	MetaBlocks []MetaBlock `json:"block,omitempty"`
	Directive  string      `json:"directive"`
}

// 倒数第三层的Block
type InnerBlock struct {
	Line         int           `json:"line"`
	Args         []string      `json:"args"`
	Directive    string        `json:"directive"`
	InmostBlocks []InmostBlock `json:"block,omitempty"`
}

// 最外面的Block, Parsed的子元素
type Block struct {
	Line        int          `json:"line"`
	Args        []string     `json:"args"`
	Directive   string       `json:"directive"`
	InnerBlocks []InnerBlock `json:"block,omitempty"`
	Includes    []string     `json:"includes,omitempty"`
}

// Parsed结构体
type Parsed struct {
	Line      int      `json:"line"`
	Args      []string `json:"args"`
	Blocks    []Block  `json:"block,omitempty"`
	Directive string   `json:"directive"`
}

// Config结构体
type Config struct {
	Status string        `json:"status"`
	Errors []interface{} `json:"errors"`
	Parsed []Parsed      `json:"parsed"`
	File   string        `json:"file"`
}

// Nginx结构体
type NginxConfig struct {
	Status  string        `json:"status"`
	Errors  []interface{} `json:"errors"`
	Configs []Config      `json:"config"`
}

func (ng *NginxConfig) Marshal() ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	data, err := json.Marshal(ng)
	return data, err
}

func (ng *NginxConfig) Unmarshal(data []byte) error {
	return jsoniter.Unmarshal(data, ng)
}

func (ng *NginxConfig) UnmarshalFromFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("ioutil ReadFile error: err=%s", err)
		return err
	}

	return ng.Unmarshal(data)
}
