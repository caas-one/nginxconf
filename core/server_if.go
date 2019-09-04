package core

import (
	"strings"
)

// determine whether it is a Server module ProxyPass directive
func isServerIfProxyPassDirective(directive string) bool {
	if isEqualString(directive, ServerProxyPassDirective) {
		return true
	}
	return false
}

// determine whether it is a Server module SendFile directive
func isServerIfSendFileDirective(directive string) bool {
	if isEqualString(directive, ServerSendFileDirective) {
		return true
	}
	return false
}

// determine whether it is a Server module Root directive
func isServerIfRootDirective(directive string) bool {
	if isEqualString(directive, ServerRootDirective) {
		return true
	}
	return false
}

// determine whether it is a Server module Rewrite directive
func isServerIfRewriteDirective(directive string) bool {
	if isEqualString(directive, ServerRewriteDirective) {
		return true
	}
	return false
}

// determine whether it is a Server module Return directive
func isServerIfReturnDirective(directive string) bool {
	if isEqualString(directive, ServerReturnDirective) {
		return true
	}
	return false
}

// determine whether it is a Server module Set directive
func isServerIfSetDirective(directive string) bool {
	if isEqualString(directive, ServerSetDirecrive) {
		return true
	}
	return false
}

// determine whether it is a Server module LimitRateAfter directive
func isServerIfLimitRateAfterDirective(directive string) bool {
	if isEqualString(directive, ServerLimitRateAfterDirective) {
		return true
	}
	return false
}

// process If conditions
func processServerIfCondition(args []string) string {
	ss := processArgsArray(args)
	return trimNewline(ss)
}

// ProcessServerIfBlocks process server if blocks
func ProcessServerIfBlocks(innerBlock InnerBlock) (*ServerIfBlock, error) {
	ifBlock := NewServerIfBlock()
	args := processQuote(innerBlock.Args)
	// If condition
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
