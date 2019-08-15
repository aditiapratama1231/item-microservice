package endpoint

import (
	"context"

	payload "bitbucket.org/qasir-id/supplier-user-service/pkg/request/payload"

	"bitbucket.org/qasir-id/supplier-user-service/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

func MakeLoginEndpoint(srv service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.LoginRequest)
		d, err := srv.LoginUser(ctx, req)
		return d, err
	}
}
