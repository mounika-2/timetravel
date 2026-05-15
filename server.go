package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rainbowmga/timetravel/api"
	"github.com/rainbowmga/timetravel/database"
	svc "github.com/rainbowmga/timetravel/service"
)

// logError logs all non-nil errors
func logError(err error) {
	if err != nil {
		log.Printf("error: %v", err)
	}
}

func main() {
	router := mux.NewRouter()

	//service := service.NewInMemoryRecordService()

	db, err := database.NewSQLiteDB("records.db")
	if err != nil {
		log.Fatal(err)
	}

	service := svc.NewSQLiteRecordService(db)
	api := api.NewAPI(&service)

	apiRoute := router.PathPrefix("/api/v1").Subrouter()
	apiV2Route := router.PathPrefix("/api/v2").Subrouter()
	api.CreateRoutes(apiV2Route)
	apiRoute.Path("/health").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		logError(err)
	})
	api.CreateRoutes(apiRoute)

	address := "127.0.0.1:8000"
	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("listening on %s", address)
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	srv.Handler = corsHandler(router)

	log.Fatal(srv.ListenAndServe())
}
