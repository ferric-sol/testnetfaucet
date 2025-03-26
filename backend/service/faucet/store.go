package faucet

import (
	"database/sql"
	"fmt"
	"solana-faucet-api/config"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// Check if IP has requested SOL in the last 24 hours
func (s *Store) HasRequested(ipAddress string) (bool, error) {
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM faucet_requests WHERE ip_address = ? AND created_at >= NOW() - INTERVAL 1 DAY", ipAddress).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Insert a new faucet request
func (s *Store) RequestSOL(ipAddress, walletAddress string) (string, error) {
	// Ensure IP hasn't already requested
	hasRequested, err := s.HasRequested(ipAddress)
	if err != nil {
		return "", err
	}
	if hasRequested {
		return "", fmt.Errorf("you can only request SOL once every 24 hours")
	}

	// Send SOL transaction
	txHash, err := SendSolanaTransaction(walletAddress)
	if err != nil {
		return "", err
	}

	// Insert new request
	_, err = s.db.Exec("INSERT INTO faucet_requests (ip_address, wallet_address, lamports, tx_hash) VALUES (?,?,?,?)",
		ipAddress,
		walletAddress,
		config.Envs.SOLTransactionLamports,
		txHash.String())
	if err != nil {
		return "", err
	}

	fmt.Println("Transaction successful! TX:", txHash.String())
	return txHash.String(), nil
}
