package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"solana-faucet-api/config"
	"solana-faucet-api/types"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, map[string]string{"error": err.Error()})
}

// VerifyRecaptchaToken sends the token to Google's reCAPTCHA verification API
func VerifyRecaptchaToken(token, ip string) bool {
	secretKey := config.Envs.RecaptchaSecretKey
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

// Helper function to format duration in a user-friendly way
// func FormatTimeDuration(d time.Duration) string {
// 	d = d.Round(time.Minute)
// 	hours := d / time.Hour
// 	d -= hours * time.Hour
// 	minutes := d / time.Minute

// 	if hours > 0 && minutes > 0 {
// 		return fmt.Sprintf("%d hours and %d minutes", hours, minutes)
// 	} else if hours > 0 {
// 		return fmt.Sprintf("%d hours", hours)
// 	}
// 	return fmt.Sprintf("%d minutes", minutes)
// }
