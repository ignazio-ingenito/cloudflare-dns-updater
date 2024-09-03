package web

import (
	"dnsupdater/api"
	"dnsupdater/views"
	"net/http"

	"github.com/a-h/templ"
	"gorm.io/gorm"
)

type PublicApiResponse struct {
	Message string
	Status  int
}

func SetupRoutes(sql *gorm.DB) {
	component := views.Home("Cloudflare DNS Updater", "Cloudflare DNS Updater")
	http.Handle("/", templ.Handler(component))
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	json.NewEncoder(w).Encode(PublicApiResponse{
	// 		Status:  http.StatusOK,
	// 		Message: "Hello world",
	// 	})
	// })

	http.HandleFunc("/api/ip/public", func(w http.ResponseWriter, r *http.Request) {
		api.ApiIpPublicGet(w, r, sql)
	})

	http.HandleFunc("/api/ip/update", func(w http.ResponseWriter, r *http.Request) {
		api.ApiIpUpdateGet(w, r, sql)
	})
}
