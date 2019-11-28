package requests

import (
	"context"
	"encoding/json"
	"net/http"
)

type UppercaseRequest struct {
	Str string `json:"str"`
}

type CountRequest struct {
	Str string `json:"str"`
}

func DecodeUppercaseRequest(_ context.Context, req *http.Request) (interface{}, error) {
	request := UppercaseRequest{}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeCountRequest(_ context.Context, req *http.Request) (interface{}, error) {
	request := CountRequest{}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
