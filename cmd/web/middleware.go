package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Writing to console from middleware")
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	crsfHandler := nosurf.New(next)
	crsfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,                 // javascript'in cookie'ye erişimini engeller
		Path:     "/",                  // cookie tüm path'lerde geçerli olur
		Secure:   false,                // http üzerinden çalışır (geliştirme için) (production'da true olmalı)
		SameSite: http.SameSiteLaxMode, // Farklı sitelerden gelen isteklerde cookie gönderilmez
	})
	return crsfHandler
}
