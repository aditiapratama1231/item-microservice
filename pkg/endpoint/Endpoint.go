package endpoint

import "github.com/go-kit/kit/endpoint"

type Endpoints struct {
	GetItemsEndpoint   endpoint.Endpoint
	CreateItemEndpoint endpoint.Endpoint
	UpdateItemEndpoint endpoint.Endpoint
	ShowItemEndpoint   endpoint.Endpoint
	DeleteItemEndpoint endpoint.Endpoint
}
