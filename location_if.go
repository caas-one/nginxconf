package nginx

// 判断是否是Location结构中If里的proxy_pass指令
func isLocationIfProxyPassDirective(directive string) bool {
	if isEqualString(directive, LocationProxyPassDirective) {
		return true
	}
	return false
}

// 判断是否是Location结构中If里的sendfile指令
func isLocationIfSendFileDirective(directive string) bool {
	if isEqualString(directive, LocationSendFileDirective) {
		return true
	}
	return false
}

// 判断是否是Location结构中If里的root指令
func isLocationIfRootDirective(directive string) bool {
	if isEqualString(directive, LocationRootDirective) {
		return true
	}
	return false
}

// 判断是否是Location结构中rewrite里的root指令
func isLocationRewriteRootDirective(directive string) bool {
	if isEqualString(directive, LocationRewriteDirective) {
		return true
	}
	return false
}

// 判断是否是location结构体中的if指令
func isLocationIfDirective(directive string) bool {
	if isEqualString(directive, LocationIfDirective) {
		return true
	}
	return false
}

// 判断是否是location if结构体中的add_header指令
func isLocationIfAddHeaderDirective(directive string) bool {
	if isEqualString(directive, AddHeaderDirective) {
		return true
	}
	return false
}

// 处理If里面的判断条件
func processLocationIfCondition(args []string) string {
	ss := processArgsArray(args)
	return trimNewline(ss)
}

func processLocationIfDirective(block InmostBlock) (*LocationIfBlock, error) {
	ifBlock := NewLocationIfBlock()

	// 判断If语句里是否需要加引号
	args := processQuote(block.Args)
	// If的条件condition
	ifBlock.Condition = processLocationIfCondition(args)
	for _, metaBlock := range block.MetaBlocks {
		// proxy_pass
		if isLocationIfProxyPassDirective(metaBlock.Directive) {
			ifBlock.ProxyPass = processArgsArray(metaBlock.Args)
			continue
		}

		// root
		if isLocationIfRootDirective(metaBlock.Directive) {
			ifBlock.Root = processArgsArray(metaBlock.Args)
			continue
		}

		// sendfile
		if isLocationIfSendFileDirective(metaBlock.Directive) {
			ifBlock.SendFile = processArgsArray(metaBlock.Args)
			continue
		}

		// rewrite
		if isLocationRewriteRootDirective(metaBlock.Directive) {
			// 判断rewrite指令里是否需要加引号
			args := processQuote(metaBlock.Args)
			ifBlock.Rewrite = append(ifBlock.Rewrite, processArgsArray(args))
		}

		// set
		if isLocationSetDirective(metaBlock.Directive) {
			ifBlock.Set = append(ifBlock.Set, processArgsArray(metaBlock.Args))
		}

		// add_header
		if isLocationIfAddHeaderDirective(metaBlock.Directive) {
			// 判断add_header指令里是否需要加引号
			for i := range metaBlock.Args {
				metaBlock.Args[i] = nonescapeQuotation(metaBlock.Args[i])
			}
			args := processQuote(metaBlock.Args)
			ifBlock.AddHeader = append(ifBlock.AddHeader, processArgsArray(args))
		}
	}

	return ifBlock, nil
}
