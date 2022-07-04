package middleware

import (
	"fmt"
	"net"
	"net/http"

	"github.com/sirupsen/logrus"
)

func RequestLogging(log bool, logger *logrus.Logger) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				ip = ""
			}

			if log == true {
				var msg = fmt.Sprintf(
					"\tReceived request (from '%s'): \"%s %v\"",
					ip, r.Method, r.URL)
				logger.Info(msg)
			}

			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
