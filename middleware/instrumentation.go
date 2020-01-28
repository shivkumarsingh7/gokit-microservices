package middleware

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/shivkumarsingh7/gokit-microservices/services"
)

type instrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           services.StringServices
}

func InstrumentingMiddleware(requestCount metrics.Counter, requestLatency, countResult metrics.Histogram) Middleware {
	return func(next services.StringServices) services.StringServices {
		return &instrumentingMiddleware{
			RequestCount:   requestCount,
			RequestLatency: requestLatency,
			CountResult:    countResult,
			Next:           next,
		}
	}
}

func (mw instrumentingMiddleware) UpperCase(s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "uppercase", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.UpperCase(s)
	return
}

func (mw instrumentingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "count", "error", "false"}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		mw.CountResult.Observe(float64(n))
	}(time.Now())

	n = mw.Next.Count(s)
	return
}
