// Package micro is a go-micro handler. It handles requests from go-micro apps.
package micro

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/joncalhoun/qson"
	"github.com/micro/go-api/handler"
	proto "github.com/micro/go-api/internal/proto"
	"github.com/micro/go-api/resolver"
	"github.com/micro/go-api/resolver/micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"github.com/micro/util/go/lib/ctx"
)

const (
	Handler = "micro"
)

type microHandler struct {
	Options  handler.Options
	Resolver resolver.Resolver
}

func (h *microHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ep, err := h.Resolver.Resolve(r)
	if err != nil {
		// we have no way of routing the request
		er := errors.InternalServerError("go.micro.api", "no route found")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte(er.Error()))
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// get content type
	ct := r.Header.Get("Content-Type")
	// create context
	cx := ctx.FromRequest(r)

	service := ep.Name
	method := ep.Method

	// get payload
	br, err := requestPayloadFromRequest(r)
	if err != nil {
		e := errors.InternalServerError("go.micro.api", err.Error())
		http.Error(w, e.Error(), 500)
		return
	}

	// Strip charset from Content-Type (like `application/json; charset=UTF-8`)
	if idx := strings.IndexRune(ct, ';'); idx >= 0 {
		ct = ct[:idx]
	}

	// client
	c := h.Options.Service.Client()

	switch ct {
	case "application/json":
		// response content type
		w.Header().Set("Content-Type", "application/json")

		var request json.RawMessage
		// if the extracted payload isn't empty lets use it
		if len(br) > 0 {
			request = json.RawMessage(br)
		}

		// create request/response
		var response json.RawMessage
		req := c.NewRequest(
			service,
			method,
			&request,
			client.WithContentType("application/json"),
		)

		// make the call
		if err := c.Call(cx, req, &response); err != nil {
			ce := errors.Parse(err.Error())
			switch ce.Code {
			case 0:
				// assuming it's totally screwed
				ce.Code = 500
				ce.Id = "go.micro.api"
				ce.Status = http.StatusText(500)
				ce.Detail = "error during request: " + ce.Detail
				w.WriteHeader(500)
			default:
				w.WriteHeader(int(ce.Code))
			}
			w.Write([]byte(ce.Error()))
			return
		}

		b, _ := response.MarshalJSON()
		w.Header().Set("Content-Length", strconv.Itoa(len(b)))
		w.Write(b)
	case "application/proto", "application/protobuf":
		request := &proto.Message{}
		// if the extracted payload isn't empty lets use it
		if len(br) > 0 {
			request = proto.NewMessage(br)
		}

		// create request/response
		response := &proto.Message{}
		req := c.NewRequest(service, method, request)

		// make the call
		if err := c.Call(cx, req, response); err != nil {
			ce := errors.Parse(err.Error())
			switch ce.Code {
			case 0:
				// assuming it's totally screwed
				ce.Code = 500
				ce.Id = "go.micro.api"
				ce.Status = http.StatusText(500)
				ce.Detail = "error during request: " + ce.Detail
				w.WriteHeader(500)
			default:
				w.WriteHeader(int(ce.Code))
			}

			// response content type
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(ce.Error()))
			return
		}

		b, _ := response.Marshal()
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		w.Header().Set("Content-Length", strconv.Itoa(len(b)))
		w.Write(b)
	default:
		http.Error(w, "Unsupported content-type", 500)
		return
	}
}

func (rh *microHandler) String() string {
	return "micro"
}

func NewHandler(opts ...handler.Option) handler.Handler {
	options := handler.NewOptions(opts...)
	return &microHandler{
		Options:  options,
		Resolver: micro.NewResolver(),
	}
}

// requestPayloadFromRequest takes a *http.Request.
// If the request is a GET the query string parameters are extracted and marshaled to JSON and the raw bytes are returned.
// If the request method is a POST the request body is read and returned
func requestPayloadFromRequest(r *http.Request) ([]byte, error) {
	switch r.Method {
	case "GET":
		if len(r.URL.RawQuery) > 0 {
			return qson.ToJSON(r.URL.RawQuery)
		}
	case "PATCH", "POST":
		return ioutil.ReadAll(r.Body)
	}

	return []byte{}, nil
}
