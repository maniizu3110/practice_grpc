package service

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"gitlab.com/techschool/pcbook/gen/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrAlreadyExists is returned when a record with the same ID already exists in the store
var ErrAlreadyExists = errors.New("record already exists")

//LaptopServer is the server that provides laptop services
type LaptopServer struct {
	Store LaptopStore
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{
		Store:store,
	}
}

func (server *LaptopServer) CreateLaptop(ctx context.Context,req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s", laptop.Id)
	if len(laptop.Id) > 0 {
		//check if it's valid uuid
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not ainvalid UUID:%v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "cannot generate a new laptop ID:%v", err)
		}
		laptop.Id = id.String()

	}
	//save the laptop to db
	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists

		}
		return nil, status.Errorf(code, "cannot save laptop to the store:%v", err)
	}

	log.Printf("saved laptop with id: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil

}
