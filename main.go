package main

import (
	"fmt"
	"gochi_api/routers"
	"net/http"
)

func main() {
	r := routers.NewRoutes()

	splash := http.FileServer(http.Dir("./template"))
	r.Handle("/", http.StripPrefix("/", splash))

	fmt.Println("api listening : http://localhost:8200")
	http.ListenAndServe(":8200", r)
}
