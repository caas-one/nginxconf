package nginxconf

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

func importcurrentDir() (*python.PyObject, error) {
	m := python.PyImport_ImportModule("sys")
	if m == nil {
		return nil, errors.New("import sys error")
	}
	return m, nil
}

// pyParse parse nginxconf  by python crossplane library
func pyParse(filePath string) (string, error) {
	m, err := importcurrentDir()
	if err != nil {
		return "", err
	}
	// import python  crossplane.parser
	m = python.PyImport_ImportModule("crossplane.parser")
	if m != nil {
		callFunc := m.GetAttrString("parse_to_json")
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
// crossplane cmd must in $PATH if use the function
func cmdNginxConf(filePath string) (string, error) {
	args := []string{"parse", filePath}
	execCmd := exec.Command("crossplane", args...)
	out, err := execCmd.CombinedOutput()
	return string(out), err
}
