package nginxconf_test

import (
	"testing"

	"github.com/caas-one/nginxconf"
	"github.com/caas-one/nginxconf/core"
)

func newTestGlobal() *core.Global {
	g := core.NewGlobal()
	g.Events.AcceptMutex = "off"
	g.Events.AcceptMutexDelay = "500ms"
	g.Events.Use = "epoll"
	g.Events.DebugConnection = make([]string, 0)
	g.Events.WorkerConnections = "65534"
	g.Daemon = "on"
	g.MasterProcess = "on"
	g.Pid = "logs/nginx.pid"
	return g
}

var wantGlobal = `
daemon on;
master_process on;
pid logs/nginx.pid;
events {
	accept_mutex off;
	accept_mutex_delay 500ms;
	use epoll;
	worker_connections 65534;
}
`

func newTestHttp() *core.Http {
	return &core.Http{
		DefaultType:               "application/octet-stream",
		ClientBodyBufferSize:      "16k",
		LargeClientHeaderBuffers:  "4 64k",
		ServerNamesHashBucketSize: "128",
		Gzip:                      "off",
		GzipMinLength:             "1232",
	}
}

var wantHttp = `http {
	default_type application/octet-stream;
	large_client_header_buffers 4 64k;
	server_names_hash_bucket_size 128;
	gzip off;
	gzip_min_length 1232;
	client_body_buffer_size 16k;
	
}
	`

func newTestServer() *core.Server {
	return &core.Server{
		Listen:    "80",
		Name:      "eg.com",
		ErrorPage: "500 502 503 504 /index.html",
		Locations: []core.Location{
			core.Location{
				Path:                 "/",
				ClientBodyBufferSize: "128k",
				ClientMaxBodySize:    "10m",
				ProxyRedirect:        "off",
				IfBlocks: []core.LocationIfBlock{
					core.LocationIfBlock{
						Condition: "$request_uri ~* /zgt/",
						ProxyPass: "http://eg.com/ortal_http",
					},
				},
			},
			core.Location{
				Path: "= /index.html",
				Root: "/eg/500",
			},
		},
	}
}

var wantServer = `server {
	server_name eg.com;
	listen 80;
	error_page 500 502 503 504 /index.html;
	
	location  / {
		client_body_buffer_size 128k;
		client_max_body_size 10m;
		proxy_redirect off;
		if ($request_uri ~* /zgt/) {
			proxy_pass http://eg.com/ortal_http;
			
		}
	}
	location  = /index.html {
		root  /eg/500;
	}
}`

func newTestUpstream() *core.Upstream {
	return &core.Upstream{
		Name:   "eg_http",
		Method: "ip_hash",
		Servers: []core.UpstreamServer{
			core.UpstreamServer{
				Address: "127.0.0.1:8080",
				Params: core.UpstreamServerParameters{
					Weight:      "99",
					MaxFails:    "3",
					FailTimeout: "10s",
				},
			},
		},
	}
}

var wantUpstream = `upstream eg_http {
	ip_hash;
	server 127.0.0.1:8080 weight=99 max_fails=3 fail_timeout=10s;
}`

func newTestDomainGroup() *core.DomainGroup {
	return &core.DomainGroup{
		Servers: []core.Server{
			core.Server{
				Listen:    "80",
				Name:      "eg.com",
				ErrorPage: "500 502 503 504 /index.html",
				Locations: []core.Location{
					core.Location{
						Path:                 "/",
						ClientBodyBufferSize: "128k",
						ClientMaxBodySize:    "10m",
						ProxyRedirect:        "off",
						IfBlocks: []core.LocationIfBlock{
							core.LocationIfBlock{
								Condition: "$request_uri ~* /zgt/",
								ProxyPass: "http://eg.com/ortal_http",
							},
						},
					},
					core.Location{
						Path: "= /index.html",
						Root: "/eg/500",
					},
				},
			},
			core.Server{
				Listen:    "81",
				Name:      "eg.com",
				ErrorPage: "500 502 503 504 /index.html",
				Locations: []core.Location{
					core.Location{
						Path:                 "/",
						ClientBodyBufferSize: "128k",
						ClientMaxBodySize:    "10m",
						ProxyRedirect:        "off",
						IfBlocks: []core.LocationIfBlock{
							core.LocationIfBlock{
								Condition: "$request_uri ~* /zgt/",
								ProxyPass: "http://eg.com/ortal_http",
							},
						},
					},
					core.Location{
						Path: "= /index.html",
						Root: "/eg/500",
					},
				},
			},
		},
		Upstreams: []core.Upstream{
			core.Upstream{
				Name:   "eg_http",
				Method: "ip_hash",
				Servers: []core.UpstreamServer{
					core.UpstreamServer{
						Address: "127.0.0.1:8080",
						Params: core.UpstreamServerParameters{
							Weight:      "99",
							MaxFails:    "3",
							FailTimeout: "10s",
						},
					},
				},
			},
			core.Upstream{
				Name:   "eg_http2",
				Method: "ip_hash",
				Servers: []core.UpstreamServer{
					core.UpstreamServer{
						Address: "127.0.0.1:8081",
						Params: core.UpstreamServerParameters{
							Weight:      "1",
							MaxFails:    "3",
							FailTimeout: "10s",
						},
					},
				},
			},
		},
	}
}

var wantDomainGroup = `upstream eg_http {
	ip_hash;
	server 127.0.0.1:8080 weight=99 max_fails=3 fail_timeout=10s;
}
upstream eg_http2 {
	ip_hash;
	server 127.0.0.1:8081 weight=1 max_fails=3 fail_timeout=10s;
}
server {
	server_name eg.com;
	listen 80;
	error_page 500 502 503 504 /index.html;
	
	location  / {
		client_body_buffer_size 128k;
		client_max_body_size 10m;
		proxy_redirect off;
		if ($request_uri ~* /zgt/) {
			proxy_pass http://eg.com/ortal_http;
			
		}
	}
	location  = /index.html {
		root  /eg/500;
	}
}
server {
	server_name eg.com;
	listen 81;
	error_page 500 502 503 504 /index.html;
	
	location  / {
		client_body_buffer_size 128k;
		client_max_body_size 10m;
		proxy_redirect off;
		if ($request_uri ~* /zgt/) {
			proxy_pass http://eg.com/ortal_http;
			
		}
	}
	location  = /index.html {
		root  /eg/500;
	}
}
`

func TestRenderToString(t *testing.T) {
	type args struct {
		tb core.TemplateBlock
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "renderGlobal",
			args: args{
				tb: newTestGlobal(),
			},
			want: wantGlobal,
		},

		{
			name: "renderHttp",
			args: args{
				tb: newTestHttp(),
			},
			want: wantHttp,
		},
		{
			name: "renderServer",
			args: args{
				tb: newTestServer(),
			},
			want: wantServer,
		},
		{
			name: "renderUpstream",
			args: args{
				tb: newTestUpstream(),
			},
			want: wantUpstream,
		},
		{
			name: "renderDomainGroup",
			args: args{
				tb: newTestDomainGroup(),
			},
			want: wantDomainGroup,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nginxconf.RenderToString(tt.args.tb)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenderToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RenderToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
