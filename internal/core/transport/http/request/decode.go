package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var requestValidator = validator.New()

type Validator interface {
	Validate() error
}

// DecodeAndValidateRequest recives pointer on struct (as an writer interface)
// and fill it in place
func DecodeAndValidateRequest(r *http.Request, destDTO any) error {

	if err := json.NewDecoder(r.Body).Decode(&destDTO); err != nil {
		return fmt.Errorf("decode json: %w", err)
	}

	var err error

	if d, ok := destDTO.(Validator); ok {
		err = d.Validate()
	} else {
		err = requestValidator.Struct(destDTO)
	}

	if err != nil {
		return fmt.Errorf(" request DTO validation: %w", err)
	}

	return nil
}
