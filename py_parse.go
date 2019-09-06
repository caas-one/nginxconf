package nginx

import (
	"errors"
	"fmt"

	"github.com/sbinet/go-python"
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

func importcurrentDir() (*python.PyObject, error) {
	m := python.PyImport_ImportModule("sys")
	if m == nil {
		return nil, errors.New("import sys error")
	}
	path := m.GetAttrString("path")
	if path != nil {
		currentDir := python.PyString_FromString("")
		if err := python.PyList_Insert(path, 0, currentDir); err != nil {
			return nil, err
		}
		return m, nil
	}
	return nil, errors.New("python get path error")
}

// pyParse parse nginxconf  by python crossplane library
func pyParse(filePath string) (string, error) {
	m, err := importcurrentDir()
	if err != nil {
		return "", err
	}
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

// CMDNginxConf parse nginx conf by crossplane parse directive
// func CMDNginxConf(filePath string) (string, error) {
// 	args := []string{"parse", filePath}
// 	execCmd := exec.Command("crossplane", args...)
// 	out, err := execCmd.CombinedOutput()
// 	return string(out), err
// }
