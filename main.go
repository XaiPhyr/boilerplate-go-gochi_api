package main

import (
	"fmt"
	"gochi_api/routers"
	"gochi_api/utils"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	r := routers.NewRoutes()
	cfg := utils.InitConfig()

	log.SetFlags(log.Llongfile | log.LstdFlags)

	rootFile := cfg.Frontend.Source
	_, err := os.Stat(rootFile + "/index.html")

	if !os.IsNotExist(err) {
		FileServer(r, rootFile)
	}

	fmt.Println()
	log.Printf("-> Local:   http://localhost:8200")
	log.Printf("-> Version: %s", cfg.Env)
	fmt.Println()

	http.ListenAndServe(":8200", r)
}

func FileServer(r chi.Router, root string) {
	fs := http.FileServer(http.Dir(root))
	r.Handle("/", http.StripPrefix("/", fs))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
			return
		}

		fs.ServeHTTP(w, r)
	})
}
