package nginxconf_test

import (
	"fmt"

	"github.com/caas-one/nginxconf"
	"github.com/caas-one/nginxconf/core"
)

func Example_nginxconf_Parse() {
	path := "./nginx.conf"
	global, err := nginxconf.Parse(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(global)
}

func Example_nginxconf_RenderToString() {
	path := "./nginx.conf"
	global, err := nginxconf.Parse(path)
	if err != nil {
		panic(err)
	}
	result, err := nginxconf.RenderToString(global)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func Example_RenderServer() {
	server := &core.Server{
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
	result, err := nginxconf.RenderToString(server)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

// Example_RenderDomainGroup single doamin render
func Example_RenderDomainGroup() {
	data, err := nginxconf.RenderToString(newTestDomainGroup())
	if err != nil {
		panic(err)
	}
	fmt.Printf(data)
}
