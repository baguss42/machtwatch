package middleware

import (
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type HandleWithError func(http.ResponseWriter, *http.Request) (int, error)

var (
	logger   *zap.Logger
)

func Load() {
	l, _ := zap.NewProduction()
	logger = l
}

func Apply(handle HandleWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		httpCode, err := handle(w, r)

		elapsed := time.Since(start).Seconds() * 1000
		elapsedStr := strconv.FormatFloat(elapsed, 'f', -1, 64)

		if err != nil {
			logger.Error(err.Error(),
				zap.String("url", r.URL.String()),
				zap.String("path", r.URL.Path),
				zap.String("method", r.Method),
				zap.String("duration", elapsedStr),
				zap.Int("http_status", httpCode),
			)
		} else {
			logger.Info("everything ok",
				zap.String("url", r.URL.String()),
				zap.String("path", r.URL.Path),
				zap.String("method", r.Method),
				zap.String("duration", elapsedStr),
				zap.Int("http_status", httpCode),
			)
		}
	}
}

// None is empty middleware, for testing usage
func None(handle HandleWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = handle(w, r)
	}
}
