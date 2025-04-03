package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"solana-faucet-api/configs"
	"solana-faucet-api/types"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

// ParseJSON now works with *gin.Context
func ParseJSON(c *gin.Context, payload any) error {
	if err := c.ShouldBindJSON(payload); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}
	return nil
}

// WriteJson now works with *gin.Context
func WriteJson(c *gin.Context, status int, v any) {
	c.JSON(status, v)
}

// WriteError now works with *gin.Context
func WriteError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{"error": err.Error()})
}

// VerifyRecaptchaToken sends the token to Google's reCAPTCHA verification API
func VerifyRecaptchaToken(token, ip string) bool {
	secretKey := configs.Envs.RecaptchaSecretKey
	if secretKey == "" {
		fmt.Println("RECAPTCHA_SECRET_KEY is not set")
		return false
	}

	// Construct request to Google reCAPTCHA API
	recaptchaURL := "https://www.google.com/recaptcha/api/siteverify"
	resp, err := http.PostForm(recaptchaURL, url.Values{
		"secret":   {secretKey},
		"response": {token},
		"remoteip": {ip},
	})
	if err != nil {
		fmt.Println("Error verifying reCAPTCHA:", err)
		return false
	}
	defer resp.Body.Close()

	// Decode response
	var result types.RecaptchaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding reCAPTCHA response:", err)
		return false
	}

	if !result.Success {
		fmt.Println("reCAPTCHA verification failed:", result.ErrorCodes)
	}

	return result.Success
}
