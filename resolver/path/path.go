// Package path resolves using http path
package path

import (
	"errors"
	"net/http"
	"strings"
)

type Resolver struct{}

func (r *Resolver) Resolve(req *http.Request) (string, error) {
	if req.URL.Path == "/" {
		return "", errors.New("unknown name")
	}
	parts := strings.Split(req.URL.Path[1:], "/")
	return parts[0], nil
}

func (r *Resolver) String() string {
	return "path"
}
