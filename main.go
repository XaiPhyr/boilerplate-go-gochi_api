package main

import (
	"gochi_api/routers"
	"log"
	"net/http"
)

func main() {
	r := routers.NewRoutes()

	splash := http.FileServer(http.Dir("./template"))
	r.Handle("/", http.StripPrefix("/", splash))

	log.Printf("-> Local:   http://localhost:8200")
	http.ListenAndServe(":8200", r)
}
