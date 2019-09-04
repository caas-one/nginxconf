package core

// determine whether it is a Location module If field proxy_pass directive
func isLocationIfProxyPassDirective(directive string) bool {
	if isEqualString(directive, LocationProxyPassDirective) {
		return true
	}
	return false
}

// determine whether it is a Location module If field sendfile directive
func isLocationIfSendFileDirective(directive string) bool {
	if isEqualString(directive, LocationSendFileDirective) {
		return true
	}
	return false
}

// determine whether it is a Location module If field root directive
func isLocationIfRootDirective(directive string) bool {
	if isEqualString(directive, LocationRootDirective) {
		return true
	}
	return false
}

// determine whether it is a Location module rewrite field root directive
func isLocationRewriteRootDirective(directive string) bool {
	if isEqualString(directive, LocationRewriteDirective) {
		return true
	}
	return false
}

// determine whether it is a location module If directive
func isLocationIfDirective(directive string) bool {
	if isEqualString(directive, LocationIfDirective) {
		return true
	}
	return false
}

// determine whether it is a location module If field add_header directive
func isLocationIfAddHeaderDirective(directive string) bool {
	if isEqualString(directive, AddHeaderDirective) {
		return true
	}
	return false
}

// process If conditions
func processLocationIfCondition(args []string) string {
	ss := processArgsArray(args)
	return trimNewline(ss)
}

func processLocationIfDirective(block InmostBlock) (*LocationIfBlock, error) {
	ifBlock := NewLocationIfBlock()
	args := processQuote(block.Args)
	// If condition need to add quotation marks
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
			// rewrite directive need to add quotation marks
			args := processQuote(metaBlock.Args)
			ifBlock.Rewrite = append(ifBlock.Rewrite, processArgsArray(args))
		}

		// set
		if isLocationSetDirective(metaBlock.Directive) {
			ifBlock.Set = append(ifBlock.Set, processArgsArray(metaBlock.Args))
		}

		// add_header
		if isLocationIfAddHeaderDirective(metaBlock.Directive) {
			// add_header directive need to add quotation marks
			for i := range metaBlock.Args {
				metaBlock.Args[i] = nonescapeQuotation(metaBlock.Args[i])
			}
			args := processQuote(metaBlock.Args)
			ifBlock.AddHeader = append(ifBlock.AddHeader, processArgsArray(args))
		}
	}

	return ifBlock, nil
}
