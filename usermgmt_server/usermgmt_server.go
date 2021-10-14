package main

import (
	//inbuilt modules
	"context"
	"log"
	"math/rand"
	"net"

	//
	pb "grpc_test/usermgmt"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	userList *pb.UserList
}

func (server *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)

	log.Printf("server listening at %v", lis.Addr())

	return s.Serve(lis)
}

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		userList: &pb.UserList{},
	}
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Recieved: %v", in.GetName())
	var user_id int32 = int32(rand.Intn(1000))
	createdUser := &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}
	s.userList.Users = append(s.userList.Users, createdUser)
	return createdUser, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	return s.userList, nil
}

func main() {
	var userMgmtServer *UserManagementServer = NewUserManagementServer()
	if err := userMgmtServer.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
