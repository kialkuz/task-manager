package httputil

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DecodeJSON[T any](r *http.Request) (*T, error) {
	var v T

	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	return &v, nil
}
