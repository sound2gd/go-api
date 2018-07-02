package router

import (
	"net/http"
	"strings"

	"github.com/micro/go-api/resolver/vpath"
)

// default resolver for legacy purposes
type defaultResolver struct {
	namespace string
	resolver  *vpath.Resolver
}

func (r *defaultResolver) Resolve(req *http.Request) (string, error) {
	name, err := r.resolver.Resolve(req)
	if err != nil {
		return "", err
	}
	return strings.Join([]string{r.namespace, name}, "."), nil
}

func (r *defaultResolver) String() string {
	return "default"
}
