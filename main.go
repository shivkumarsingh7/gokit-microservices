package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/shivkumarsingh7/gokit-microservices/endpoints"
	"github.com/shivkumarsingh7/gokit-microservices/middleware"
	"github.com/shivkumarsingh7/gokit-microservices/requests"
	"github.com/shivkumarsingh7/gokit-microservices/responses"
	"github.com/shivkumarsingh7/gokit-microservices/services"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	var svc services.StringServices
	svc = services.StrService{}
	svc = middleware.LoggingMiddleware(logger)(svc)
	uppercaseHandler := httptransport.NewServer(
		endpoints.MakeUppercaseEndpoint(svc),
		requests.DecodeUppercaseRequest,
		responses.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		endpoints.MakeCountEndpoint(svc),
		requests.DecodeCountRequest,
		responses.EncodeResponse,
	)

	fmt.Println("Intro to Microservices")

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.ListenAndServe(":8080", nil)
}
