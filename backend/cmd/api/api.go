package api

import (
	"database/sql"
	"log"
	"net/http"
	"solana-faucet-api/service/faucet"

	"github.com/gorilla/mux"
)

type APIserver struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIserver {
	return &APIserver{
		addr: addr,
		db:   db,
	}
}

func (s *APIserver) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	faucetStore := faucet.NewStore(s.db)
	faucetHandler := faucet.NewHandler(faucetStore)
	faucetHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	// Resolved CORS error when running locally
	// return http.ListenAndServe(s.addr,
	// 	handlers.CORS(
	// 		handlers.AllowedOrigins([]string{"*"}),
	// 		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	// 		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	// 	)(router))

	return http.ListenAndServe(s.addr, router)
}
