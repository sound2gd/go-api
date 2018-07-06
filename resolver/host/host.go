// Package host resolves using http host
package host

import (
	"net/http"

	"github.com/micro/go-api/resolver"
)

type Resolver struct{}

func (r *Resolver) Resolve(req *http.Request) (*resolver.Endpoint, error) {
	return &resolver.Endpoint{
		Name: req.Host,
	}, nil
}

func (r *Resolver) String() string {
	return "host"
}
