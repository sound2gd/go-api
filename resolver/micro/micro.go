// Package micro provides a micro rpc resolver which prefixes a namespace
package micro

import (
	"net/http"

	"github.com/micro/go-api/resolver"
)

// default resolver for legacy purposes
// it uses proxy routing to resolve names
// /foo becomes namespace.foo
// /v1/foo becomes namespace.v1.foo
type microResolver struct {
	opts resolver.Options
}

func (r *microResolver) Resolve(req *http.Request) (*resolver.Endpoint, error) {
	name := req.Header.Get("X-Micro-Target")
	method := req.Header.Get("X-Micro-Method")

	// internal micro client request
	if len(name) > 0 && len(method) > 0 {
		return &resolver.Endpoint{
			Name:   name,
			Method: method,
		}, nil
	}

	switch r.opts.Handler {
	// internal handlers
	case "meta", "api", "rpc", "micro":
		name, method = apiRoute(r.opts.Namespace, req.URL.Path)
	default:
		method = req.Method
		name = proxyRoute(r.opts.Namespace, req.URL.Path)
	}

	return &resolver.Endpoint{
		Name:   name,
		Method: method,
	}, nil
}

func (r *microResolver) String() string {
	return "micro"
}

// NewResolver creates a new micro resolver
func NewResolver(opts ...resolver.Option) resolver.Resolver {
	return &microResolver{
		opts: resolver.NewOptions(opts...),
	}
}
