package transport

import (
	"context"

	"github.com/aditiapratama1231/item-microservice/pkg/endpoint"

	pb "github.com/aditiapratama1231/adit-microservice/proto/item"
	transport "github.com/aditiapratama1231/item-microservice/pkg/request/grpc"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServcer struct {
	showItemDetail grpctransport.Handler
}

func (s *grpcServcer) ShowItemDetail(ctx context.Context, r *pb.ShowItemDetailRequest) (*pb.ShowItemDetailResponse, error) {
	_, resp, err := s.showItemDetail.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ShowItemDetailResponse), nil
}

func ItemGRPCServer(ctx context.Context, endpoints endpoint.Endpoints) pb.ItemsServer {
	return &grpcServcer{
		showItemDetail: grpctransport.NewServer(
			endpoints.ShowItemEndpoint,
			transport.DecodeGRPCShowItemDetailRequest,
			transport.EncodeGRPCShowItemDetailResponse,
		),
	}
}
