package nginxconf

import (
	"encoding/json"
	"fmt"

	"github.com/caas-one/nginxconf/core"
)

// Parse parse nginx conf form file path
func Parse(filePath string) (*core.Global, error) {
	ngconf, err := ParseToConfigs(filePath)
	if err != nil {
		return nil, err
	}
	// use index = 0
	conf := ngconf.Configs[0]
	globalConf, err := core.ProcessGlobal(&conf)
	if err != nil {
		return nil, fmt.Errorf("ProcessMain failed. err:%v", err)
	}
	return globalConf, nil
}

// ParseToJSON parse nginx conf to json
func ParseToJSON(filePath string) (string, error) {
	parseResult, err := pyParse(filePath)
	if err != nil {
		return "", fmt.Errorf("python parse failed. err:%v", err)
	}
	return parseResult, nil
}

// ParseToConfigs parse nginx conf to crossplane configs
func ParseToConfigs(filePath string) (*core.CrossplaneOut, error) {
	parseResult, err := ParseToJSON(filePath)
	if err != nil {
		return nil, err
	}
	ngconf := &core.CrossplaneOut{}
	if err := json.Unmarshal([]byte(parseResult), ngconf); err != nil {
		return nil, err
	}
	if len(ngconf.Configs) <= 0 {
		err := fmt.Errorf("parse nginx conf failed. configs is empty")
		return nil, err
	}
	if ngconf.Status != "ok" {
		confjson, _ := json.Marshal(ngconf)
		return nil, fmt.Errorf("parse failed. info: %s", string(confjson))
	}
	return ngconf, nil
}
