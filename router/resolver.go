package router

import (
	"net/http"
	"strings"

	"github.com/micro/go-api/resolver/vpath"
)

// default resolver for legacy purposes
type resolver struct {
	namespace string
	resolver  *vpath.Resolver
}

func (r *resolver) Resolve(req *http.Request) (string, error) {
	name, err := r.resolver.Resolve(req)
	if err != nil {
		return "", err
	}
	return strings.Join([]string{r.namespace, name}, "."), nil
}

func (r *resolver) String() string {
	return "default"
}
