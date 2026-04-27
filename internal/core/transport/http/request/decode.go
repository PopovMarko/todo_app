package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var requestValidator = validator.New()

// DecodeAndValidateRequest recives pointer on struct (as an writer interface)
// and fill it in place
func DecodeAndValidateRequest(r *http.Request, destDTO any) error {
	if err := json.NewDecoder(r.Body).Decode(&destDTO); err != nil {
		return fmt.Errorf("decode json: %w", err)
	}
	// Validate DTO if error returns use err assertion if needed
	if err := requestValidator.Struct(destDTO); err != nil {
		return fmt.Errorf(" DTO validation: %w", err)
	}
	return nil
}
