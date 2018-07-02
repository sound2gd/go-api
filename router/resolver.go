package router

import (
	"net/http"

	"github.com/micro/go-api/resolver/vpath"
)

// default resolver for legacy purposes
type defaultResolver struct {
	namespace string
	resolver  *vpath.Resolver
}

func (r *defaultResolver) Resolve(req *http.Request) (string, error) {
	name := proxyRoute(r.namespace, req.URL.Path)
	return name, nil
}

func (r *defaultResolver) String() string {
	return "default"
}
