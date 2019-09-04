package nginx

import (
	"io/ioutil"

	"github.com/json-iterator/go"
)

// MetaBlock meta Block, the innermost Block
type MetaBlock struct {
	Line      int      `json:"line"`
	Args      []string `json:"args"`
	Directive string   `json:"directive"`
}

// InmostBlock block on the penultimate level
type InmostBlock struct {
	Line       int         `json:"line"`
	Args       []string    `json:"args"`
	MetaBlocks []MetaBlock `json:"block,omitempty"`
	Directive  string      `json:"directive"`
}

// InnerBlock Block on the penultimate third tier
type InnerBlock struct {
	Line         int           `json:"line"`
	Args         []string      `json:"args"`
	Directive    string        `json:"directive"`
	InmostBlocks []InmostBlock `json:"block,omitempty"`
}

// Block Outermost Block, a child element of Parsed
type Block struct {
	Line        int          `json:"line"`
	Args        []string     `json:"args"`
	Directive   string       `json:"directive"`
	InnerBlocks []InnerBlock `json:"block,omitempty"`
	Includes    []string     `json:"includes,omitempty"`
}

// Parsed block
type Parsed struct {
	Line      int      `json:"line"`
	Args      []string `json:"args"`
	Blocks    []Block  `json:"block,omitempty"`
	Directive string   `json:"directive"`
}

// Config python crossplan library
type Config struct {
	Status string        `json:"status"`
	Errors []interface{} `json:"errors"`
	Parsed []Parsed      `json:"parsed"`
	File   string        `json:"file"`
}

// CrossplaneOut parse nginx conf result
type CrossplaneOut struct {
	Status  string        `json:"status"`
	Errors  []interface{} `json:"errors"`
	Configs []Config      `json:"config"`
}

// Marshal json marshal
func (ng *CrossplaneOut) Marshal() ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	data, err := json.Marshal(ng)
	return data, err
}

// Unmarshal json unmarshal
func (ng *CrossplaneOut) Unmarshal(data []byte) error {
	return jsoniter.Unmarshal(data, ng)
}

// UnmarshalFromFile json unmarshal from file
func (ng *CrossplaneOut) UnmarshalFromFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return ng.Unmarshal(data)
}
