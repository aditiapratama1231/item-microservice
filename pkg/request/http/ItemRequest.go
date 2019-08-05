package request

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	payload "github.com/aditiapratama1231/item-microservice/pkg/request/payload"
)

func DecodeGetItemResponse(ctx context.Context, r *http.Request) (interface{}, error) {
	return payload.GetItemResponse{}, nil
}

func DecodeCreateItemRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req payload.CreateItemRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	return req, err
}

func DecodeUpdateItemRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)

	req := payload.UpdateItemRequest{
		ID: qs["id"],
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func DecodeShowItemRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)

	req := payload.ShowItemRequest{
		ID: qs["id"],
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func DecodeDeleteItemRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)

	req := payload.DeleteItemRequest{
		ID: qs["id"],
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
