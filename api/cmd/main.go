package main

import (
	"database/sql"
	"log"
	"solana-faucet-api/cmd/api"
	"solana-faucet-api/configs"
	"solana-faucet-api/db"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// Initialize MySQL Storage
	database, err := db.NewMySQLStorage(mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	initStorage(database.DB)

	// Initialize and start the API server
	server := api.NewAPIServer(":8080", database.DB)
	if err := server.Run(); err != nil {
		log.Fatal("Server error:", err)
	}
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal("Database connection failed:", err)
	}
	log.Println("DB-server: Successfully connected!")
}
