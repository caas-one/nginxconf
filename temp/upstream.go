package temp

import (
	"fmt"
	"text/template"
)

// ServerTpl upstream template
var (
	upstreamTpl           *template.Template
	annotationUpstreamTpl *template.Template
)

func init() {
	var err error
	upstreamTpl, err = template.New("upstream_tpl").Parse(UpstreamTemp)
	if err != nil {
		panic(fmt.Errorf("init upstream_tpl failed err:%v", err))
	}
	annotationUpstreamTpl, err = template.New("annotation_upstream_tpl").Parse(AnnotationUpstreamTemp)
	if err != nil {
		panic(fmt.Errorf("init annotation_upstream_tpl failed err:%v", err))
	}
}

// GetUpstreamTpl get upstream template
func GetUpstreamTpl() *template.Template {
	return upstreamTpl
}

// GetAnnotationUpstreamTpl get annotation-upstream template
func GetAnnotationUpstreamTpl() *template.Template {
	return annotationUpstreamTpl
}

// upstream module
const (
	UpstreamTemp = `upstream {{.Name}} {
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
	AnnotationUpstreamTemp = `# upstream {{.Name}} {
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
