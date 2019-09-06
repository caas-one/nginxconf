package nginxconf

import (
	"github.com/caas-one/nginxconf/core"
)

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
