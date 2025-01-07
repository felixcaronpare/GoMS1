package handler

import (
	"context"
	"errors"
	"log"

	"github.com/felixcaronpare/GoMS1/internal/models"
	interfaces "github.com/felixcaronpare/GoMS1/pkg/v1"
	pb "github.com/felixcaronpare/GoMS1/proto"
	"google.golang.org/grpc"
)

type AccountServStruct struct {
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
    log.Printf("Received CreateAccountRequest: Name=%s, Email=%s, Password=%s", req.GetName(), req.GetEmail(), req.GetPassword())

	data := srv.transformAccountRPC(req)
	if data.Email == "" || data.Password == "" || data.Name == "" {
		return &pb.AccountProfileResponse{}, errors.New("All fields must be filled")
	}

	account, err :=  srv.usecase.Create(data)
	if err != nil {
		return &pb.AccountProfileResponse{}, err
	}
	return srv.transformAccountModel(account), nil
}


//=========== UTILS ===========

func (srv *AccountServStruct) transformAccountRPC(req *pb.CreateAccountRequest) models.Account {
    return models.Account{
        Name:     req.GetName(),
        Email:    req.GetEmail(),
        Password: req.GetPassword(),
    }
}
func (srv *AccountServStruct) transformAccountModel(account models.Account) *pb.AccountProfileResponse {
	return &pb.AccountProfileResponse{Id: string(account.ID), Name: account.Name, Email: account.Email, Password: account.Password}
  }