package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

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
	// adding instrumenting
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	svc = middleware.InstrumentingMiddleware(requestCount, requestLatency, countResult)(svc)

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
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
