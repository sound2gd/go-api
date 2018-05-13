// Package resolver resolves a http request to an endpoint
package resolver

import (
	"net/http"
)

// Resolver resolves requests to endpoints
type Resolver interface {
	Resolve(r *http.Request) (string, error)
	String() string
}
