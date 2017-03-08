package api

import (
	"strings"

	"github.com/micro/go-micro/server"
)

const (
	// Default defines the default handler
	Default Handler = "api"
	// serves api.Request and api.Response
	Api Handler = "api"
	// services an RPC request/response
	Rpc Handler = "rpc"
	// proxies a http request
	Proxy Handler = "proxy"
)

// Handler defines the type of handler uses by the micro api
type Handler string

// Endpoint is a mapping between an RPC method and HTTP endpoint
type Endpoint struct {
	// RPC Method e.g. Greeter.Hello
	Name string
	// Description e.g what's this endpoint for
	Description string
	// API Handler e.g rpc, proxy
	Handler Handler
	// HTTP Host e.g example.com
	Host []string
	// HTTP Methods e.g GET, POST
	Method []string
	// HTTP Path e.g /greeter
	Path []string
}

// Encode encodes an endpoint to endpoint metadata
func Encode(e *Endpoint) map[string]string {
	if e == nil {
		return nil
	}

	return map[string]string{
		"endpoint":    e.Name,
		"description": e.Description,
		"method":      strings.Join(e.Method, ","),
		"path":        strings.Join(e.Path, ","),
		"host":        strings.Join(e.Host, ","),
		"handler":     string(e.Handler),
	}
}

// Decode decodes endpoint metadata into an endpoint
func Decode(e map[string]string) *Endpoint {
	if e == nil {
		return nil
	}

	return &Endpoint{
		Name:        e["endpoint"],
		Description: e["description"],
		Method:      strings.Split(e["method"], ","),
		Path:        strings.Split(e["path"], ","),
		Host:        strings.Split(e["host"], ","),
		Handler:     Handler(e["handler"]),
	}
}

// WithEndpoint returns a server.HandlerOption with endpoint metadata set
// usage:
// proto.Register(server, handler, api.WithEndpoint(&Endpoint{Name: "Greeter.Hello", Path: []string{"/greeter"}}))
func WithEndpoint(e *Endpoint) server.HandlerOption {
	return server.EndpointMetadata(e.Name, Encode(e))
}
