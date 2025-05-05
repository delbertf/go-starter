package assets

import (
	"embed"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

//go:embed all:dist
var Assets embed.FS

// ServeAssets mount the embedded assets to an HTTP server
func Mount(r chi.Router) {
	r.Route("/dist", func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r)
			})
		})

		if inProduction() {
			r.Handle("/*", staticProd())
		} else {
			r.Handle("/*", disableCache(staticDev()))
		}
		//r.Handle("/*", http.FileServer(http.FS(Assets)))
	})
}

func disableCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func staticDev() http.Handler {
	return http.StripPrefix("/dist/", http.FileServerFS(os.DirFS("internal/assets/dist")))
}

func staticProd() http.Handler {
	return http.StripPrefix("/dist/", http.FileServerFS(Assets))
}

func inProduction() bool {
	return os.Getenv("APP_ENV") == "production"
}
