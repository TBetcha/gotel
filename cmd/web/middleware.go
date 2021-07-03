package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

// basic middleware for chi I use it the same way  "mux.use(middleware name)
// it takes a http handler and returns a http handler with a lambda inside which does what I want and calls next

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page, brody")
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRF protection to all post requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves the session on ev                                                                                                                                                                                                                                                             lllllllllllllllllllllllllllllllllllllllrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
