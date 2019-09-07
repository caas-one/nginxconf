package nginxconf_test

import (
	"reflect"
	"testing"

	"github.com/caas-one/nginxconf"
	"github.com/caas-one/nginxconf/core"
)

func NewTestGlobal() *core.Global {
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

func TestParse(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *core.Global
		wantErr bool
	}{
		{
			name:    "parsetoglobal",
			args:    args{filePath: "./nginx.conf"},
			want:    NewTestGlobal(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nginxconf.Parse(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

var jsonStr = `{"status": "ok", "errors": [], "config": [{"status": "ok", "errors": [], "parsed": [{"line": 1, "args": ["on"], "directive": "daemon"}, {"line": 2, "args": ["on"], "directive": "master_process"}, {"line": 3, "args": ["logs/nginx.pid"], "directive": "pid"}, {"line": 4, "args": [], "block": [{"line": 5, "args": ["off"], "directive": "accept_mutex"}, {"line": 6, "args": ["500ms"], "directive": "accept_mutex_delay"}, {"line": 7, "args": ["epoll"], "directive": "use"}, {"line": 8, "args": ["65534"], "directive": "worker_connections"}], "directive": "events"}], "file": "./nginx.conf"}]}`

func TestParseToJSON(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "parsetojson",
			args:    args{filePath: "./nginx.conf"},
			want:    jsonStr,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nginxconf.ParseToJSON(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

var testOut = &core.CrossplaneOut{
	Status: "ok",
	Errors: make([]interface{}, 0),
	Configs: []core.Config{
		core.Config{
			Status: "ok",
			Errors: make([]interface{}, 0),
			Parsed: []core.Parsed{
				core.Parsed{
					Line:      1,
					Args:      []string{"on"},
					Directive: "daemon",
				},
				core.Parsed{
					Line:      2,
					Args:      []string{"on"},
					Directive: "master_process",
				},
				core.Parsed{
					Line:      3,
					Args:      []string{"logs/nginx.pid"},
					Directive: "pid",
				},
				core.Parsed{
					Line: 4,
					Args: make([]string, 0),
					Blocks: []core.Block{
						core.Block{
							Line:      5,
							Args:      []string{"off"},
							Directive: "accept_mutex",
						},
						core.Block{
							Line:      6,
							Args:      []string{"500ms"},
							Directive: "accept_mutex_delay",
						},
						core.Block{
							Line:      7,
							Args:      []string{"epoll"},
							Directive: "use",
						},
						core.Block{
							Line:      8,
							Args:      []string{"65534"},
							Directive: "worker_connections",
						},
					},
					Directive: "events",
				},
			},
			File: "./nginx.conf",
		},
	},
}

func TestParseToConfigs(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *core.CrossplaneOut
		wantErr bool
	}{
		{
			name:    "parsetoout",
			args:    args{filePath: "./nginx.conf"},
			want:    testOut,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nginxconf.ParseToConfigs(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToConfigs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseToConfigs() = %v, want %v", got, tt.want)
			}
		})
	}
}
