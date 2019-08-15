package request

import (
	"context"
	"encoding/json"
	"net/http"

	payload "bitbucket.org/qasir-id/supplier-user-service/pkg/request/payload"
)

//EncodeResponse encode out outgoing response
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func EncodeInstropectionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(payload.TokenInstropectionResponse)

	if res.StatusCode == 403 {
		w.WriteHeader(http.StatusForbidden)
	}

	if res.StatusCode == 400 {
		w.WriteHeader(http.StatusBadRequest)
	}

	return json.NewEncoder(w).Encode(res)
}
