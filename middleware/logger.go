package middleware

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/shivkumarsingh7/gokit-microservices/services"
)

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next services.StringServices) services.StringServices {
		return &loggingMiddleware{
			Logger: logger,
			Next:   next,
		}
	}
}

type loggingMiddleware struct {
	Logger log.Logger
	Next   services.StringServices
}

func (mw loggingMiddleware) UpperCase(s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.UpperCase(s)
	return
}

func (mw loggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.Next.Count(s)
	return
}
