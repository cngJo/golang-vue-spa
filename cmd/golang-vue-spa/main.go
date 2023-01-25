package main

import (
	"io/fs"
	"net/http"

	"github.com/cngJo/golang-vuejs-spa/internal/frontend"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router(),
	}

	srv.ListenAndServe()
}

func router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

	r.Get("/", indexHandler)

	staticFS, _ := fs.Sub(frontend.StaticFiles, "dist")
	httpFS := http.FileServer(http.FS(staticFS))
	r.Handle("/assets/*", httpFS)

	r.Get("/api/v1/greeting", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"message\": \"hello from golang\"}"))
	})

	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		rawFile, _ := frontend.StaticFiles.ReadFile("dist/favicon.ico")
		w.Write(rawFile)
		return
	}
	rawFile, _ := frontend.StaticFiles.ReadFile("dist/index.html")
	w.Write(rawFile)
}
