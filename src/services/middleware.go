package services

import (
	"gomscode/src/logger"
	"net/http"
	"time"
)

// Middleware exported
func Middleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		//		logger.LogOut(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
		f(w, r)
		requesterIP := r.RemoteAddr
		logger.LogOutInfo(r.Method, r.RequestURI, requesterIP, r.Host, time.Since(start).String(), r.UserAgent())
	}
}
