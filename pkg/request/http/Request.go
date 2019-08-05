package request

import (
	"context"
	"encoding/json"
	"net/http"
)

//EncodeResponse encode out outgoing response
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
