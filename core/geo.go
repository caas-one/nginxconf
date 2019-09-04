package core

import (
	"errors"
)

// determine whether it is a geo directive
func isGeoDirective(directive string) bool {
	if isEqualString(directive, GeoDirective) {
		return true
	}
	return false
}

// determine whether it is a ranges directive
func isGeoRangesDirective(directive string) bool {
	if isEqualString(directive, GeoRangesDirective) {
		return true
	}
	return false
}

// determine whether it is a delete directive
func isGeoDeleteDirective(directive string) bool {
	if isEqualString(directive, GeoDeleteDirective) {
		return true
	}
	return false
}

// determine whether it is a default directive
func isGeoDefaultDirective(directive string) bool {
	if isEqualString(directive, GeoDefaultDirective) {
		return true
	}
	return false
}

// determine whether it is a include directive
func isGeoIncludeDirective(directive string) bool {
	if isEqualString(directive, GeoIncludeDirective) {
		return true
	}
	return false
}

// determine whether it is a include directive
func isGeoProxyDirective(directive string) bool {
	if isEqualString(directive, GeoProxyDirective) {
		return true
	}
	return false
}

// determine whether it is a include directive
func isGeoProxyRecursiveDirective(directive string) bool {
	if isEqualString(directive, GeoProxyRecursiveDirective) {
		return true
	}
	return false
}

// GeoNotListDirective geo not array directive
var GeoNotListDirective = []string{GeoRangesDirective, GeoDeleteDirective, GeoDefaultDirective, GeoIncludeDirective, GeoProxyDirective, GeoProxyRecursiveDirective}

// determine whether it is a ip list directive, eg: 123.125.62.248-123.125.62.248 1; 1.85.49.29-1.85.49.29 1;
func isGeoListDirective(directive string) bool {
	for _, str := range GeoNotListDirective {
		if isEqualString(directive, str) {
			return false
		}
	}
	return true
}

// ProcessGeo process geo directive
func ProcessGeo(block *Block) (*Geo, error) {
	if !isGeoDirective(block.Directive) {
		return nil, errors.New("Not geo directive")
	}

	geo := NewGeo()
	geo.Name = processArgsArray(block.Args)
	for _, innerBlock := range block.InnerBlocks {
		// process ranges directive
		if isGeoRangesDirective(innerBlock.Directive) {
			geo.Ranges = innerBlock.Directive
			continue
		}

		// process delete directive
		if isGeoDeleteDirective(innerBlock.Directive) {
			geo.Delete = processArgsArray(innerBlock.Args)
			continue
		}

		// process default directive
		if isGeoDefaultDirective(innerBlock.Directive) {
			geo.Default = processArgsArray(innerBlock.Args)
			continue
		}

		// process include directive
		if isGeoIncludeDirective(innerBlock.Directive) {
			geo.Include = processArgsArray(innerBlock.Args)
			continue
		}

		// process proxy directive
		if isGeoProxyDirective(innerBlock.Directive) {
			geo.Proxy = processArgsArray(innerBlock.Args)
			continue
		}

		// process proxy_recursive
		if isGeoProxyRecursiveDirective(innerBlock.Directive) {
			geo.ProxyRecursive = processArgsArray(innerBlock.Args)
			continue
		}

		// process ip list directive
		if isGeoListDirective(innerBlock.Directive) {
			// eg: 123.125.62.248-123.125.62.248 1
			str := innerBlock.Directive + " " + processArgsArray(innerBlock.Args)
			geo.List = append(geo.List, str)
			continue
		}
	}
	return geo, nil
}
