package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func NoSurf(next http.Handler) http.Handler {
	crsfHandler := nosurf.New(next)
	crsfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,                 // javascript'in cookie'ye erişimini engeller
		Path:     "/",                  // cookie tüm path'lerde geçerli olur
		Secure:   app.InProduction,     // http üzerinden çalışır (geliştirme için) (production'da true olmalı)
		SameSite: http.SameSiteLaxMode, // Farklı sitelerden gelen isteklerde cookie gönderilmez
	})
	return crsfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
