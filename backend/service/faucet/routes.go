package faucet

import (
	"fmt"
	"net/http"
	"solana-faucet-api/config"
	"solana-faucet-api/types"
	"solana-faucet-api/utils"

	"github.com/gagliardetto/solana-go"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.FaucetRequestStore
}

func NewHandler(store types.FaucetRequestStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/request", h.handleRequestSOL).Methods("POST")
	router.HandleFunc("/balance", h.handleGetBalance).Methods("GET")
}

func (h *Handler) handleRequestSOL(w http.ResponseWriter, r *http.Request) {
	ipAddress := GetClientIP(r) // Get user's IP

	// Get JSON payload
	var payload types.FaucetRequestPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// Validate Solana address format
	if !IsValidSolanaAddress(payload.WalletAddress) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid Solana wallet address"))
		return
	}

	// Verify reCAPTCHA token
	isValidRecaptcha := utils.VerifyRecaptchaToken(payload.RecaptchaToken, ipAddress)
	if !isValidRecaptcha {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("reCAPTCHA verification failed"))
		return
	}

	// Process the request
	txHash, err := h.store.RequestSOL(ipAddress, payload.WalletAddress)
	if err != nil {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("%v", err))
		return
	}

	// Success Response
	utils.WriteJson(w, http.StatusOK, map[string]string{"txHash": txHash})
}

func (h *Handler) handleGetBalance(w http.ResponseWriter, r *http.Request) {
	walletAddress := config.Envs.FundingWalletPublicKey
	if walletAddress == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("wallet address is required"))
		return
	}

	// Validate Solana address format
	if _, err := solana.PublicKeyFromBase58(walletAddress); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid Solana wallet address"))
		return
	}

	balance, err := GetWalletBalance(walletAddress)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// utils.WriteJson(w, http.StatusOK, map[string]uint64{"balance": balance})
	utils.WriteJson(w, http.StatusOK, map[string]float64{
		"balance": float64(balance) / 1_000_000_000, // Convert lamports to SOL
	})
}
