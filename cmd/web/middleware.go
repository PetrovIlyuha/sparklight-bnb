package main

import (
	"fmt"
	"net/http"
)

func WriteLogsToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("Page was hit")
		next.ServeHTTP(rw, req)
	})
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}