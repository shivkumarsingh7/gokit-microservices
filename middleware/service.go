package middleware

import "github.com/shivkumarsingh7/gokit-microservices/services"

// Middleware describes a service middleware.
type Middleware func(service services.StringServices) services.StringServices
