// Package vpath resolves using http path and recognised versioned urls
package path

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
)

type Resolver struct{}

var (
	re = regexp.MustCompile("^v[0-9]+$")
)

func (r *Resolver) Resolve(req *http.Request) (string, error) {
	if req.URL.Path == "/" {
		return "", errors.New("unknown name")
	}

	parts := strings.Split(req.URL.Path[1:], "/")

	if len(parts) == 1 {
		return parts[0], nil
	}

	// /v1/foo
	if re.MatchString(parts[0]) {
		return parts[1], nil
	}

	return parts[0], nil
}

func (r *Resolver) String() string {
	return "path"
}
