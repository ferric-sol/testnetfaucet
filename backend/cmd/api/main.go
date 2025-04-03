package api

import (
	"database/sql"
	"log"
	"solana-faucet-api/services/faucet"

	// "solana-faucet-api/service/faucet"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	r := gin.Default()

	// CORS Middleware (Replaces handlers.CORS)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Register routes
	api := r.Group("/api/v1")
	faucetStore := faucet.NewStore(s.db)
	faucetHandler := faucet.NewHandler(faucetStore)
	faucetHandler.RegisterRoutes(api)

	log.Println("Listening on", s.addr)
	return r.Run(s.addr)
}
