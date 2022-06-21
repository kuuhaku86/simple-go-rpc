package main

import (
	"errors"
	"context"

	pb "github.com/kuuhaku86/simple-go-rpc/dataserver"
)

type DataServer struct {
	pb.UnimplementedDataServerServer
}

func (server *DataServer) GetData(ctx context.Context, request *pb.Request) (*pb.Data, error) {
	for _, data := range Datas {
		if request.Key == data.Key {
			return &data, nil
		}
	}

	return nil, errors.New("Data not found")
}

func newServer() *DataServer {
	s := &DataServer{}
	return s
}