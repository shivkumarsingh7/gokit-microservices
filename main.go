package main

import (
	"fmt"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/shivkumarsingh7/gokit-microservices/endpoints"
	"github.com/shivkumarsingh7/gokit-microservices/requests"
	"github.com/shivkumarsingh7/gokit-microservices/responses"
	"github.com/shivkumarsingh7/gokit-microservices/services"
)

func main() {
	svc := services.StrService{}

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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
