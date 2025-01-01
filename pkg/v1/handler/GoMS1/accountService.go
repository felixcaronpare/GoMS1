package handler

import (
	"context"
	"errors"

	"github.com/felixcaronpare/GoMS1/internal/models"
	interface "github.com/felixcaronpare/GoMS1/pkg/v1"
	pb "github.com/felixcaronpare/GoMS1/proto"
	"google.golang.org/grpc"
)

type AccountServStruct {
	usecase interfaces.UsecaseInterface
	pb.UnimplementedAccountServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase_p interfaces.UsecaseInterface) {
	accountGrpc := &AccountServStruct{usecase: usecase_p}
	pb.RegisterAccountServiceServer(grpcServer, accountGrpc)
}

//=========== CRUD ===========

//Create
func (srv *AccountServStruct) Create(ctx context.Context, req *pb.CreateAccountRequest) (*pb.AccountProfileResponse, error) {
	data := srv.transformAccountRPC(req)
	if data.email == "" || data.password ="" {
		return &pb.AccountProfileResponse{}, errors.New("All fields must be filled")
	}

	account, err :=  srv.usecase.Create(data)
	if err != nil {
		return &pb.AccountProfileResponse{}, err
	}
	return srv.transformAccountRPC(account), nil
}