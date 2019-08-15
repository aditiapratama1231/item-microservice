package cmd

import (
	"bitbucket.org/qasir-id/supplier-user-service/database"
	"bitbucket.org/qasir-id/supplier-user-service/pkg/endpoint"
	"bitbucket.org/qasir-id/supplier-user-service/pkg/service"
)

var (
	db = database.DBInit()

	userSrv   = service.NewUserService(db)
	Endpoints = endpoint.Endpoints{
		LoginEndpoint: endpoint.MakeLoginEndpoint(userSrv),
	}
)
