package nginx

import (
	"strings"
)

// 判断是否是Server结构体中的ProxyPass指令
func isServerIfProxyPassDirective(directive string) bool {
	if isEqualString(directive, ServerProxyPassDirective) {
		return true
	}
	return false
}

// 判断是否是Server结构体中的SendFile指令
func isServerIfSendFileDirective(directive string) bool {
	if isEqualString(directive, ServerSendFileDirective) {
		return true
	}
	return false
}

// 判断是否是Server结构体中的Root指令
func isServerIfRootDirective(directive string) bool {
	if isEqualString(directive, ServerRootDirective) {
		return true
	}
	return false
}

// 判断是否是Server结构体中的Rewrite指令
func isServerIfRewriteDirective(directive string) bool {
	if isEqualString(directive, ServerRewriteDirective) {
		return true
	}
	return false
}

// 判断是否是Server结构体中的Return指令
func isServerIfReturnDirective(directive string) bool {
	if isEqualString(directive, ServerReturnDirective) {
		return true
	}
	return false
}

// 判断是否是Server结构体中的Set指令
func isServerIfSetDirective(directive string) bool {
	if isEqualString(directive, ServerSetDirecrive) {
		return true
	}
	return false
}

// 判断是否是Server结构体中的LimitRateAfter指令
func isServerIfLimitRateAfterDirective(directive string) bool {
	if isEqualString(directive, ServerLimitRateAfterDirective) {
		return true
	}
	return false
}

// 处理If里面的判断条件
func processServerIfCondition(args []string) string {
	ss := processArgsArray(args)

	// 判断condition里面是否有换行
	return trimNewline(ss)
}

func ProcessServerIfBlocks(innerBlock InnerBlock) (*ServerIfBlock, error) {
	ifBlock := NewServerIfBlock()

	// 过滤condition是否需要加引号
	args := processQuote(innerBlock.Args)

	// If的条件condition
	ifBlock.Condition = processServerIfCondition(args)

	for _, inmost := range innerBlock.InmostBlocks {
		// ProxyPass
		if isServerIfProxyPassDirective(inmost.Directive) {
			ifBlock.ProxyPass = processArgsArray(inmost.Args)
			continue
		}

		// SendFile
		if isServerIfSendFileDirective(inmost.Directive) {
			ifBlock.SendFile = processArgsArray(inmost.Args)
			continue
		}

		// Root
		if isServerIfRootDirective(inmost.Directive) {
			ifBlock.Root = processArgsArray(inmost.Args)
			continue
		}

		// Rewrite
		if isServerIfRewriteDirective(inmost.Directive) {
			// 判断rewrite指令是否需要加引号
			args := processQuote(inmost.Args)
			ifBlock.Rewrite = append(ifBlock.Rewrite, processArgsArray(args))
			continue
		}

		// Return
		if isServerIfReturnDirective(inmost.Directive) {
			ifBlock.Return = processArgsArray(inmost.Args)
			continue
		}

		// Set
		if isServerIfSetDirective(inmost.Directive) {
			if len(inmost.Args) > 1 {
				for index := range inmost.Args {
					if strings.EqualFold(inmost.Args[index], "") {
						inmost.Args[index] = "''"
					}
				}
				ifBlock.Set = append(ifBlock.Set, processArgsArray(inmost.Args))
			}
			continue
		}

		// LimitRateAfter
		if isServerIfLimitRateAfterDirective(inmost.Directive) {
			ifBlock.LimitRateAfter = processArgsArray(inmost.Args)
			continue
		}

	}

	return ifBlock, nil
}
