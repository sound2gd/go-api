package router

import (
	"net/http"
)

// default resolver for legacy purposes
// it uses proxy routing to resolve names
// /foo becomes namespace.foo
// /v1/foo becomes namespace.v1.foo
type defaultResolver struct {
	namespace string
}

func (r *defaultResolver) Resolve(req *http.Request) (string, error) {
	name := proxyRoute(r.namespace, req.URL.Path)
	return name, nil
}

func (r *defaultResolver) String() string {
	return "default"
}
