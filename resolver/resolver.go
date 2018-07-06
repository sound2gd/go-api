// Package resolver resolves a http request to an endpoint
package resolver

import (
	"net/http"
)

// Resolver resolves requests to endpoints
type Resolver interface {
	Resolve(r *http.Request) (*Endpoint, error)
	String() string
}

// Endpoint is the endpoint for a http request
type Endpoint struct {
	// e.g greeter
	Name string
}
