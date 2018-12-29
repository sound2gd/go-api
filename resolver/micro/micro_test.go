package micro

import (
	"net/http"
	"testing"
)

func TestMicroResolver(t *testing.T) {
	r := &http.Request{
		Header: map[string][]string{
			"X-Micro-Target": {"foo"},
			"X-Micro-Method": {"Foo.Bar"},
		},
	}

	rr := NewResolver()
	ep, err := rr.Resolve(r)
	if err != nil {
		t.Fatal(err)
	}
	if ep.Name != "foo" {
		t.Fatalf("Expected endpoint name foo got: %s", ep.Name)
	}
	if ep.Method != "Foo.Bar" {
		t.Fatalf("Expected endpoint method Foo.Bar got: %s", ep.Method)
	}
}
