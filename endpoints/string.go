package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/shivkumarsingh7/gokit-microservices/requests"
	"github.com/shivkumarsingh7/gokit-microservices/responses"
	"github.com/shivkumarsingh7/gokit-microservices/services"
)

func MakeUppercaseEndpoint(svc services.StrService) endpoint.Endpoint {
	return func(_ context.Context, req interface{}) (interface{}, error) {
		r := req.(requests.UppercaseRequest)
		str, err := svc.UpperCase(r.Str)
		if err != nil {
			return responses.UppercaseResponse{str, err}, nil
		}
		return responses.UppercaseResponse{str, nil}, nil
	}
}

func MakeCountEndpoint(svc services.StrService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.CountRequest)
		len := svc.Count(req.Str)
		return responses.CountResponse{len}, nil
	}
}
