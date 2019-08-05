package request

import (
	"context"
	"fmt"
	"strconv"

	pb "github.com/aditiapratama1231/adit-microservice/proto/item"
	payload "github.com/aditiapratama1231/item-microservice/pkg/request/payload"

	"github.com/jinzhu/copier"
	"github.com/joho/godotenv"
)

func EncodeGRPCShoItemDetailRequest(_ context.Context, r interface{}) (interface{}, error) {
	return &pb.ShowItemDetailRequest{}, nil
}

func DecodeGRPCShowItemDetailRequest(ctx context.Context, r interface{}) (interface{}, error) {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

	data := r.(*pb.ShowItemDetailRequest)

	item := payload.ShowItemRequest{
		ID: strconv.FormatInt(int64(data.ItemId), 10),
	}

	return item, nil
}

func EncodeGRPCShowItemDetailResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(payload.ItemResponse)
	data := &pb.Item{}
	copier.Copy(data, resp.Data)
	return &pb.ShowItemDetailResponse{
		Data: data,
	}, nil
}
