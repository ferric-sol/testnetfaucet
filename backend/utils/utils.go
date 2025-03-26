package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

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
