package nginx

import (
	"errors"
)

// 判断是否是geo指令
func isGeoDirective(directive string) bool {
	if isEqualString(directive, GeoDirective) {
		return true
	}
	return false
}

// 判断是否是ranges指令
func isGeoRangesDirective(directive string) bool {
	if isEqualString(directive, GeoRangesDirective) {
		return true
	}
	return false
}

// 判断是否是delete指令
func isGeoDeleteDirective(directive string) bool {
	if isEqualString(directive, GeoDeleteDirective) {
		return true
	}
	return false
}

// 判断是否是default指令
func isGeoDefaultDirective(directive string) bool {
	if isEqualString(directive, GeoDefaultDirective) {
		return true
	}
	return false
}

// 判断是否是include指令
func isGeoIncludeDirective(directive string) bool {
	if isEqualString(directive, GeoIncludeDirective) {
		return true
	}
	return false
}

// 判断是否是include指令
func isGeoProxyDirective(directive string) bool {
	if isEqualString(directive, GeoProxyDirective) {
		return true
	}
	return false
}

// 判断是否是include指令
func isGeoProxyRecursiveDirective(directive string) bool {
	if isEqualString(directive, GeoProxyRecursiveDirective) {
		return true
	}
	return false
}

var GeoNotListDirective = []string{GeoRangesDirective, GeoDeleteDirective, GeoDefaultDirective, GeoIncludeDirective, GeoProxyDirective, GeoProxyRecursiveDirective}

// 判断是否是ip list指令，例如：123.125.62.248-123.125.62.248 1; 1.85.49.29-1.85.49.29 1;
// 用排除法判断，非ranges等指令就认为是ip list
func isGeoListDirective(directive string) bool {
	for _, str := range GeoNotListDirective {
		if isEqualString(directive, str) {
			return false
		}
	}
	return true
}

// ProcessGeo 处理geo指令
func ProcessGeo(block *Block) (*Geo, error) {
	if !isGeoDirective(block.Directive) {
		return nil, errors.New("Not geo directive")
	}

	geo := NewGeo()
	// Geo的$address $variable作为geo的Name
	geo.Name = processArgsArray(block.Args)
	for _, innerBlock := range block.InnerBlocks {
		// 处理ranges指令
		if isGeoRangesDirective(innerBlock.Directive) {
			geo.Ranges = innerBlock.Directive
			continue
		}

		// 处理delete指令
		if isGeoDeleteDirective(innerBlock.Directive) {
			geo.Delete = processArgsArray(innerBlock.Args)
			continue
		}

		// 处理default指令
		if isGeoDefaultDirective(innerBlock.Directive) {
			geo.Default = processArgsArray(innerBlock.Args)
			continue
		}

		// 处理include指令
		if isGeoIncludeDirective(innerBlock.Directive) {
			geo.Include = processArgsArray(innerBlock.Args)
			continue
		}

		// 处理proxy指令
		if isGeoProxyDirective(innerBlock.Directive) {
			geo.Proxy = processArgsArray(innerBlock.Args)
			continue
		}

		// 处理proxy_recursive
		if isGeoProxyRecursiveDirective(innerBlock.Directive) {
			geo.ProxyRecursive = processArgsArray(innerBlock.Args)
			continue
		}

		// 处理ip list指令
		if isGeoListDirective(innerBlock.Directive) {
			// 123.125.62.248-123.125.62.248 1
			str := innerBlock.Directive + " " + processArgsArray(innerBlock.Args)
			geo.List = append(geo.List, str)
			continue
		}
	}
	return geo, nil
}
