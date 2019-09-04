package nginx

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func isValidHost(arg string) bool {
	ipReg := regexp.MustCompile(IPRegex)
	if ok := ipReg.MatchString(arg); ok {
		return true
	}

	domainReg := regexp.MustCompile(DomainRegex)
	if ok := domainReg.MatchString(arg); ok {
		return true
	}

	return false
}

func isValidPort(arg string) bool {
	port, err := strconv.Atoi(arg)
	if err != nil {
		return false
	}

	if port > 0 && port < 65535 {
		return true
	}
	return false
}

func isUnixProtocol(arg string) bool {
	return strings.HasPrefix(arg, "unix:")
}

func isEqualString(str1, str2 string) bool {
	if strings.EqualFold(strings.ToLower(str1), strings.ToLower(str2)) {
		return true
	}
	return false
}

func processArgsArray(args []string) string {
	return strings.Join(args, " ")
}

// If/Location/Rewrite add quote
func processQuote(args []string) []string {
	for index, value := range args {
		if isNeededQuote(value) {
			args[index] = "\"" + value + "\""
		}
	}
	return args
}

// doc：http://nginx.org/en/docs/http/ngx_http_rewrite_module.html
func isNeededQuote(arg string) bool {
	if strings.Contains(arg, "}") || strings.Contains(arg, ";") {
		return true
	}
	return false
}

func trimNewline(arg string) string {
	if len(arg) > 0 && strings.Contains(arg, "\n") {
		return strings.Replace(arg, "\n", "", -1)
	}
	return arg
}

var (
	old, new = fmt.Sprintf("%s", "\""), fmt.Sprintf("\\%s", "\"")
)

func nonescapeQuotation(str string) string {
	if strings.Contains(str, old) {
		return strings.Replace(str, old, new, -1)
	}
	return str
}
