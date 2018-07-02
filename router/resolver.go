package router

import (
	"net/http"

	"github.com/micro/go-api"
)

// default resolver for legacy purposes
// it uses proxy routing to resolve names
// /foo becomes namespace.foo
// /v1/foo becomes namespace.v1.foo
type defaultResolver struct {
	handler   api.Handler
	namespace string
}

func (r *defaultResolver) Resolve(req *http.Request) (string, error) {
	switch r.handler {
	case api.Default, api.Api, api.Rpc:
		name, _ := apiRoute(r.namespace, req.URL.Path)
		return name, nil
	default:
		name := proxyRoute(r.namespace, req.URL.Path)
		return name, nil
	}
}

func (r *defaultResolver) String() string {
	return "default"
}
