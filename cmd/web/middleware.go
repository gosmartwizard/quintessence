package main

import (
	"net/http"
)

/*
	   fun middleware(next http.Handler) http.Handler{

		   fn := func(w http.ResponseWriter, r *http.Request) {

			   // execute the before logic of middleware

			   // call the next http.Handler

			   // execute the after logic of middleware
		   }

		   return http.HandlerFunc(fn
	   }
*/
func secureHeaders(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")

		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")

		w.Header().Set("X-Content-Type-Options", "nosniff")

		w.Header().Set("X-Frame-Options", "deny")

		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
