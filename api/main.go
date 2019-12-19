package main

import (
	"net/http"
	"newbug/routers"
)

func main() {
	router:= routers.InitRouter()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
	}

	server.ListenAndServe()
}
