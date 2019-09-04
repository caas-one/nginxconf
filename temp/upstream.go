package temp

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/caas-one/nginx-go/core"
)

// ServerTpl upstream template
var (
	UpstreamTpl           *template.Template
	AnnotationUpstreamTpl *template.Template
)

func init() {
	var err error
	UpstreamTpl, err = template.New("upstream_tpl").Parse(Upstream)
	if err != nil {
		panic(fmt.Errorf("init upstream_tpl failed err:%v", err))
	}
	AnnotationUpstreamTpl, err = template.New("annotation_upstream_tpl").Parse(AnnotationUpstream)
	if err != nil {
		panic(fmt.Errorf("init annotation_upstream_tpl failed err:%v", err))
	}
}

// UpstreamToConf use template render upstream config
func UpstreamToConf(conf core.Upstream) ([]byte, error) {
	content := new(bytes.Buffer)
	err := UpstreamTpl.Execute(content, conf)
	return content.Bytes(), err
}

// AnnotationUpstreamToConf use template render upstream config
func AnnotationUpstreamToConf(conf core.Upstream) ([]byte, error) {
	content := new(bytes.Buffer)
	err := AnnotationUpstreamTpl.Execute(content, conf)
	return content.Bytes(), err
}

// upstream module
const (
	Upstream = `upstream {{.Name}} {
	{{- if .Method}}
	{{.Method}};
	{{- end -}}
	{{range $server := .Servers}}
	server {{$server.Address}} {{- if $server.Params.Weight}} weight={{$server.Params.Weight}} {{- end -}} {{- if $server.Params.MaxConns}} max_conns={{$server.Params.MaxConns}} {{- end -}} {{- if $server.Params.MaxFails}} max_fails={{$server.Params.MaxFails}} {{- end -}} {{- if $server.Params.FailTimeout}} fail_timeout={{$server.Params.FailTimeout}} {{- end -}} {{- if $server.Params.Backup}} backup {{- end}} {{- if $server.Params.Down}} down {{- end}};
	{{- end -}}
	{{if .UpstreamKeepalive.Time}}
	keepalive {{.UpstreamKeepalive.Time}};
	{{- end}}
	{{- if .KeepaliveTimeout}}
	keepalive_timeout {{.KeepaliveTimeout}};
	{{- end }}
}`
)

// AnnotationUpstream annotation upstream
const (
	AnnotationUpstream = `# upstream {{.Name}} {
 	{{- if .Method}}
# 	{{.Method}};
 	{{- end -}}
 	{{range $server := .Servers}}
# 	server {{$server.Address}} {{- if $server.Params.Weight}} weight={{$server.Params.Weight}} {{- end -}} {{- if $server.Params.MaxConns}} max_conns={{$server.Params.MaxConns}} {{- end -}} {{- if $server.Params.MaxFails}} max_fails={{$server.Params.MaxFails}} {{- end -}} {{- if $server.Params.FailTimeout}} fail_timeout={{$server.Params.FailTimeout}} {{- end -}} {{- if $server.Params.Backup}} backup {{- end}} {{- if $server.Params.Down}} down {{- end}};
 	{{- end -}}
 	{{if .UpstreamKeepalive.Time}}
# 	keepalive {{.UpstreamKeepalive.Time}};
	{{- end}}
	{{- if .KeepaliveTimeout}}
#	keepalive_timeout {{.KeepaliveTimeout}};
	{{- end }}
# }`
)
