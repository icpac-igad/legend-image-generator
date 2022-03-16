package service

import (
	"time"

	"github.com/gocraft/web"
	log "github.com/sirupsen/logrus"
)

// loggerMiddleware is generic middleware that will log requests to log.
// extends web.LoggerMiddleware to log request methods
func loggerMiddleware(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	startTime := time.Now()

	next(rw, req)

	duration := time.Since(startTime).Nanoseconds()
	var durationUnits string
	switch {
	case duration > 2000000:
		durationUnits = "ms"
		duration /= 1000000
	case duration > 1000:
		durationUnits = "μs"
		duration /= 1000
	default:
		durationUnits = "ns"
	}

	log.Printf("%s [%d %s] %d '%s'\n", req.Method, duration, durationUnits, rw.StatusCode(), req.URL.Path)
}
