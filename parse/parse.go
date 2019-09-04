package parse

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/sbinet/go-python"
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

// NginxConf parse nginxconf  by python crossplane library
func NginxConf(filePath string) (string, error) {
	m := python.PyImport_ImportModule("sys")
	if m == nil {
		return "", errors.New("import sys error")
	}
	path := m.GetAttrString("path")
	if path != nil {
		currentDir := python.PyString_FromString("")
		python.PyList_Insert(path, 0, currentDir)
		m = python.PyImport_ImportModule("dict_to_json")
		if m != nil {
			callFunc := m.GetAttrString("dict_to_json")
			out := callFunc.CallFunction(python.PyString_FromString(filePath))
			outstr := fmt.Sprintf("%v", out)
			if outstr != "" {
				return outstr, nil
			}
			return "", errors.New("call python Module crossplane function parse  error, output is empty")
		}
		return "", errors.New("PyImport_ImportModule failed")
	}
	return "", errors.New("python get path error")
}

// CMDNginxConf parse nginx conf by crossplane parse directive
func CMDNginxConf(filePath string) (string, error) {
	args := []string{"parse", filePath}
	execCmd := exec.Command("crossplane", args...)
	out, err := execCmd.CombinedOutput()
	return string(out), err
}
