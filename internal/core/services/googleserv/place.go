package googleserv

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/domain/googledm"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type service struct {
	apiKey string
}

func NewService(
	apiKey string,
) *service {
	return &service{
		apiKey: apiKey,
	}
}

func (svc *service) Queryautocomplete(condition *googledm.QueryautocompleteRequest) (res *googledm.QueryautocompleteResponse, err error) {
	errFrom := func(err error) (*googledm.QueryautocompleteResponse, error) {
		logrus.Errorln(err)
		return nil, err
	}

	condition.Key = svc.apiKey

	queries, err := CreateQuery(condition)
	if err != nil {
		return errFrom(err)
	}

	url := "https://maps.googleapis.com/maps/api/place/queryautocomplete/json"

	if queries != "" {
		url = url + "?" + queries
	}

	req, resp := NewRequest(nil, GET, url)

	{
		err := fasthttp.Do(req, resp)
		if err != nil {
			return errFrom(err)
		}
		fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)
	}

	{
		body := resp.Body()
		if resp.StatusCode() != fasthttp.StatusOK {
			return errFrom(errors.New(string(body)))
		}

		reader := bytes.NewBuffer(body)
		err := json.NewDecoder(reader).Decode(&res)
		if err != nil {
			return errFrom(err)
		}
	}
	return res, nil
}

func (svc *service) PlaceDetails(condition *googledm.PlaceDetailsRequest) (res *googledm.PlaceDetailsResponse, err error) {
	errFrom := func(err error) (*googledm.PlaceDetailsResponse, error) {
		logrus.Errorln(err)
		return nil, err
	}

	condition.Key = svc.apiKey

	queries, err := CreateQuery(condition)
	if err != nil {
		return errFrom(err)
	}

	url := "https://maps.googleapis.com/maps/api/place/details/json"

	if queries != "" {
		url = url + "?" + queries
	}

	req, resp := NewRequest(nil, GET, url)

	{
		err := fasthttp.Do(req, resp)
		if err != nil {
			return errFrom(err)
		}
		fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)
	}

	{
		body := resp.Body()
		if resp.StatusCode() != fasthttp.StatusOK {
			return errFrom(errors.New(string(body)))
		}

		reader := bytes.NewBuffer(body)
		err := json.NewDecoder(reader).Decode(&res)
		if err != nil {
			return errFrom(err)
		}
	}
	return res, nil
}
