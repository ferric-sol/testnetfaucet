package types

import "time"

type FaucetRequestStore interface {
	RequestSOL(ipAddress, walletAddress string) (string, error)
	// HasRequested(ipAddress string) (bool, error)
	GetLastRequestTime(ipAddress string) (time.Time, error)
}

type FaucetRequestPayload struct {
	WalletAddress  string `json:"walletAddress" validate:"required"`
	RecaptchaToken string `json:"recaptchaToken" validate:"required"`
}

type FaucetRequest struct {
	ID            int       `json:"id"`
	IPAddress     string    `json:"ip_address"`
	WalletAddress string    `json:"wallet_address"`
	Amount        int       `json:"amount"`
	TXHash        string    `json:"tx_hash"`
	CreatedAt     time.Time `json:"created_at"`
}

type RecaptchaResponse struct {
	Success     bool     `json:"success"`
	ChallengeTs string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes,omitempty"`
}
