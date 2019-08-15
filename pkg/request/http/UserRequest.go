package request

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	payload "bitbucket.org/qasir-id/supplier-user-service/pkg/request/payload"
)

func DecodeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req payload.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	fmt.Println(req)
	return req, err
}
