package httppresentation

import "net/http"

func responseHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getAllMiddlewareList() (middleware []func(http.Handler) http.Handler) {
	middleware = append(middleware, responseHeaders)

	return middleware
}
