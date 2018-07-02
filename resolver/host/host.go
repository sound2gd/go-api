// Package host resolves using http host
package host

import (
	"net/http"
)

type Resolver struct{}

func (r *Resolver) Resolve(req *http.Request) (string, error) {
	return req.Host, nil
}

func (r *Resolver) String() string {
	return "host"
}
