package endpoint

import (
	"context"

	payload "github.com/aditiapratama1231/item-microservice/pkg/request/payload"

	"github.com/aditiapratama1231/item-microservice/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

func MakeGetItemsEndpoint(srv service.ItemService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		d := srv.GetItems()
		return d, nil
	}
}

func MakeCreateItemEndpoint(srv service.ItemService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.CreateItemRequest)
		d := srv.CreateItem(req)
		return d, nil
	}
}

func MakeUpdateItemEndpoint(srv service.ItemService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.UpdateItemRequest)
		d := srv.UpdateItem(req)
		return d, nil
	}
}

func MakeShowItemEndpoint(srv service.ItemService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.ShowItemRequest)
		d := srv.ShowItem(req.ID)
		return d, nil
	}
}

func MakeDeleteItemEndpoint(srv service.ItemService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.DeleteItemRequest)
		d := srv.DeleteItem(req.ID)
		return d, nil
	}
}
