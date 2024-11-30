package transport

import (
	"context"
	"e-commerce/service/kit/endpoint"
	pb "e-commerce/service/kit/transport/grpc"
)

type grpcServer struct {
	endpoints endpoint.Set
}

func NewGrpcServer(endpoints endpoint.Set) pb.RetailServiceServer {
	return &grpcServer{
		endpoints: endpoints,
	}
}

func (g grpcServer) CreateProduct(ctx context.Context, req *pb.CreateProductReq) (*pb.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (g grpcServer) UpdateProduct(ctx context.Context, req *pb.UpdateProductReq) (*pb.Response, error) {
	//TODO implement me
	panic("implement me")
}
