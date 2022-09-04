package googleserv

import (
	"github.com/google/go-querystring/query"
	"github.com/valyala/fasthttp"
)

var (
	POST            = []byte(fasthttp.MethodPost)
	GET             = []byte(fasthttp.MethodGet)
	PUT             = []byte(fasthttp.MethodPut)
	PATCH           = []byte(fasthttp.MethodPatch)
	DELETE          = []byte(fasthttp.MethodDelete)
	ApplicationJSON = []byte("application/json")
)

func NewRequest(body []byte, method []byte, url string) (*fasthttp.Request, *fasthttp.Response) {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	req.SetBody(body)
	req.Header.SetMethodBytes(method)
	req.SetRequestURIBytes([]byte(url))

	return req, resp
}

func CreateQuery(inf interface{}) (string, error) {
	v, err := query.Values(inf)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}
