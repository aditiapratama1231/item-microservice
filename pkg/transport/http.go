package transport

import (
	"context"
	"net/http"

	request "github.com/aditiapratama1231/item-microservice/pkg/request/http"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/aditiapratama1231/item-microservice/pkg/endpoint"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	item := r.PathPrefix("/api/items").Subrouter()
	item.Use(commonMiddleware)

	getItemsHandler := httptransport.NewServer(
		endpoints.GetItemsEndpoint,
		request.DecodeGetItemResponse,
		request.EncodeResponse,
	)

	createItemHandler := httptransport.NewServer(
		endpoints.CreateItemEndpoint,
		request.DecodeCreateItemRequest,
		request.EncodeResponse,
	)

	updatetItemHandler := httptransport.NewServer(
		endpoints.UpdateItemEndpoint,
		request.DecodeUpdateItemRequest,
		request.EncodeResponse,
	)

	showItemHandler := httptransport.NewServer(
		endpoints.ShowItemEndpoint,
		request.DecodeShowItemRequest,
		request.EncodeResponse,
	)

	deleteItemHandler := httptransport.NewServer(
		endpoints.DeleteItemEndpoint,
		request.DecodeDeleteItemRequest,
		request.EncodeResponse,
	)

	item.Handle("", getItemsHandler).Methods("GET")
	item.Handle("/{id}", showItemHandler).Methods("GET")
	item.Handle("/create", createItemHandler).Methods("POST")
	item.Handle("/{id}/update", updatetItemHandler).Methods("PATCH")
	item.Handle("/{id}/delete", deleteItemHandler).Methods("DELETE")

	return item
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
