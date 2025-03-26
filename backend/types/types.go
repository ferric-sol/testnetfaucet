package types

import "time"

type FaucetRequestStore interface {
	RequestSOL(ipAddress, walletAddress string) (string, error)
	HasRequested(ipAddress string) (bool, error)
}

type FaucetRequestPayload struct {
	WalletAddress string `json:"walletAddress" validate:"required"`
}

type FaucetRequest struct {
	ID            int       `json:"id"`
	IPAddress     string    `json:"ip_address"`
	WalletAddress string    `json:"wallet_address"`
	Amount        int       `json:"amount"`
	TXHash        string    `json:"tx_hash"`
	CreatedAt     time.Time `json:"created_at"`
}
