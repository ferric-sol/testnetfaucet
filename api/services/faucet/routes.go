package faucet

import (
	"fmt"
	"net/http"
	"solana-faucet-api/configs"
	"solana-faucet-api/types"
	"solana-faucet-api/utils"

	"github.com/gagliardetto/solana-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Handler struct with store dependency
type Handler struct {
	store types.FaucetRequestStore
}

// NewHandler initializes the faucet handler
func NewHandler(store types.FaucetRequestStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes sets up the Gin routes for the faucet API
func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/request", h.handleRequestSOL)
	router.GET("/balance", h.handleGetBalance)
}

// handleRequestSOL processes a SOL request
func (h *Handler) handleRequestSOL(c *gin.Context) {
	ipAddress := c.ClientIP() // Get user's IP from Gin context

	// Parse JSON payload
	var payload types.FaucetRequestPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid payload: %v", errors)})
		return
	}

	// Validate Solana address format
	if !IsValidSolanaAddress(payload.WalletAddress) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Solana wallet address"})
		return
	}

	// Verify reCAPTCHA token
	if !utils.VerifyRecaptchaToken(payload.RecaptchaToken, ipAddress) {
		c.JSON(http.StatusForbidden, gin.H{"error": "reCAPTCHA verification failed"})
		return
	}

	// Process the request
	txHash, err := h.store.RequestSOL(ipAddress, payload.WalletAddress)
	if err != nil {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
		return
	}

	// Success Response
	c.JSON(http.StatusOK, gin.H{"txHash": txHash})
}

// handleGetBalance retrieves the faucet wallet balance
func (h *Handler) handleGetBalance(c *gin.Context) {
	walletAddress := configs.Envs.FundingWalletPublicKey
	if walletAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wallet address is required"})
		return
	}

	// Validate Solana address format
	if _, err := solana.PublicKeyFromBase58(walletAddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Solana wallet address"})
		return
	}

	balance, err := GetWalletBalance(walletAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert lamports to SOL and respond
	c.JSON(http.StatusOK, gin.H{"balance": float64(balance) / 1_000_000_000})
}
