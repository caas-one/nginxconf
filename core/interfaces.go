package core

import (
	"bytes"
	"errors"

	"github.com/caas-one/nginxconf/temp"
)

// TemplateBlock template block
type TemplateBlock interface {
	Render() ([]byte, error)
}

var _ TemplateBlock = &Global{}

// Render use template reander global nginx config
func (global *Global) Render() ([]byte, error) {
	content := new(bytes.Buffer)
	if err := temp.GetGlobalTpl().Execute(content, global); err != nil {
		return nil, err
	}
	return content.Bytes(), nil
}

var _ TemplateBlock = &Http{}

// Render use http template render http nginx confi
func (http *Http) Render() ([]byte, error) {
	content := new(bytes.Buffer)
	if err := temp.GetHTTPTpl().Execute(content, http); err != nil {
		return nil, err
	}
	return content.Bytes(), nil
}

var _ TemplateBlock = &Server{}

// Render use template render server nginx config
func (server *Server) Render() ([]byte, error) {
	content := new(bytes.Buffer)
	if err := temp.GetServerTpl().Execute(content, server); err != nil {
		return nil, err
	}
	return content.Bytes(), nil
}

var _ TemplateBlock = &Upstream{}

// Render use template render upstream config
func (upstream *Upstream) Render() ([]byte, error) {
	content := new(bytes.Buffer)
	if err := temp.GetUpstreamTpl().Execute(content, upstream); err != nil {
		return nil, err
	}
	return content.Bytes(), nil
}

var _ TemplateBlock = &AnnotationUpstream{}

// Render use template render annotation upstream config
func (annotationUpstream *AnnotationUpstream) Render() ([]byte, error) {
	content := new(bytes.Buffer)
	if err := temp.GetAnnotationUpstreamTpl().Execute(content, annotationUpstream); err != nil {
		return nil, err
	}
	return content.Bytes(), nil
}

// Render render domain group
func (domainGroup DomainGroup) Render() ([]byte, error) {
	var byts bytes.Buffer
	if len(domainGroup.Upstreams) > 0 {
		for i := range domainGroup.Upstreams {
			upstream := domainGroup.Upstreams[i]
			upstreamByts, err := upstream.Render()
			if err != nil {
				return nil, err
			}
			if _, err = byts.Write(upstreamByts); err != nil {
				return nil, err
			}
			byts.WriteByte('\n')
		}
	}

	if len(domainGroup.Servers) > 0 {
		for i := range domainGroup.Servers {
			server := domainGroup.Servers[i]
			serverByts, err := server.Render()
			if err != nil {
				return nil, err
			}
			if _, err = byts.Write(serverByts); err != nil {
				return nil, err
			}
			byts.WriteByte('\n')
		}
	}

	r := byts.Bytes()
	if len(r) <= 0 {
		return nil, errors.New("data is empty")
	}
	return r, nil
}
