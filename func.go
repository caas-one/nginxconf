package nginx

import (
	"encoding/json"
	"fmt"

	"github.com/caas-one/nginx-go/core"

	"github.com/caas-one/nginx-go/python"
)

// Parse parse nginx conf form file path
func Parse(filePath string) (*core.Global, error) {
	parseResult, err := python.NginxConf(filePath)
	if err != nil {
		return nil, fmt.Errorf("python parse failed. err:%v", err)
	}
	ngconf := core.CrossplaneOut{}
	if err := json.Unmarshal([]byte(parseResult), &ngconf); err != nil {
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
	// use index = 0
	conf := ngconf.Configs[0]
	globalConf, err := core.ProcessGlobal(&conf)
	if err != nil {
		return nil, fmt.Errorf("ProcessMain failed. err:%v", err)
	}
	return globalConf, nil
}

// RenderToBytes render module to nginx conf
func RenderToBytes(tb core.TemplateBlock) ([]byte, error) {
	return tb.Render()
}

// RenderToString render module to nginx conf
func RenderToString(tb core.TemplateBlock) (string, error) {
	byts, err := tb.Render()
	if err != nil {
		return "", err
	}
	return string(byts), nil
}
