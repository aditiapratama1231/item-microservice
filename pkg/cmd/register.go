package cmd

import (
	"github.com/aditiapratama1231/item-service/database"
	"github.com/aditiapratama1231/item-service/pkg/endpoint"
	"github.com/aditiapratama1231/item-service/pkg/service"
)

var (
	db = database.DBInit()

	srvItem   = service.NewIttemService(db)
	Endpoints = endpoint.Endpoints{
		GetItemsEndpoint:   endpoint.MakeGetItemsEndpoint(srvItem),
		CreateItemEndpoint: endpoint.MakeCreateItemEndpoint(srvItem),
		UpdateItemEndpoint: endpoint.MakeUpdateItemEndpoint(srvItem),
		ShowItemEndpoint:   endpoint.MakeShowItemEndpoint(srvItem),
		DeleteItemEndpoint: endpoint.MakeDeleteItemEndpoint(srvItem),
	}
)
