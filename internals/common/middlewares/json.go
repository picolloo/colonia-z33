package middlewares

import "net/http"

func JsonMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}
