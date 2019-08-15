package transport

import (
	"context"
	"net/http"

	request "bitbucket.org/qasir-id/supplier-user-service/pkg/request/http"
	httptransport "github.com/go-kit/kit/transport/http"

	"bitbucket.org/qasir-id/supplier-user-service/pkg/endpoint"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	user := r.PathPrefix("/api").Subrouter()
	user.Use(commonMiddleware)

	loginHandler := httptransport.NewServer(
		endpoints.LoginEndpoint,
		request.DecodeLoginRequest,
		request.EncodeResponse,
	)

	tokenInstropectionHandler := httptransport.NewServer(
		endpoints.TokenInstropection,
		request.DecodeTokenInstropectionRequest,
		request.EncodeInstropectionResponse,
	)

	user.Handle("/login", loginHandler).Methods("POST")
	user.Handle("/token-instropection", tokenInstropectionHandler).Methods("GET")
	return user
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
