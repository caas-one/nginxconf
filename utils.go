package nginx

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// 判断是ip地址或者是域名
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

// 判断是否是合法的端口号
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

// 判断是否以unix://开头
func isUnixProtocol(arg string) bool {
	return strings.HasPrefix(arg, "unix:")
}

// 判断两个字符串是否相等
func isEqualString(str1, str2 string) bool {
	if strings.EqualFold(strings.ToLower(str1), strings.ToLower(str2)) {
		return true
	}
	return false
}

// 将args数组转化成一个字符串，用空格分开
func processArgsArray(args []string) string {
	return strings.Join(args, " ")
}

// If/Location/Rewrite语句块正则语句加引号
func processQuote(args []string) []string {
	for index, value := range args {
		if isNeededQuote(value) {
			args[index] = "\"" + value + "\""
		}
	}
	return args
}

// 是否包含引号和分号
// 官方文档：http://nginx.org/en/docs/http/ngx_http_rewrite_module.html
func isNeededQuote(arg string) bool {
	if strings.Contains(arg, "}") || strings.Contains(arg, ";") {
		return true
	}
	return false
}

// 加密套件过长会被crossplane加换行符，要手动去掉
func trimNewline(arg string) string {
	if len(arg) > 0 && strings.Contains(arg, "\n") {
		return strings.Replace(arg, "\n", "", -1)
	}
	return arg
}

var (
	old, new = fmt.Sprintf("%s", "\""), fmt.Sprintf("\\%s", "\"")
)

// NonescapeQuotation 把内容带"的字符串转义
func nonescapeQuotation(str string) string {
	if strings.Contains(str, old) {
		return strings.Replace(str, old, new, -1)
	}
	return str
}
