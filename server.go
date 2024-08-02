package main

import (
    "context"
    "log"
    "net"
    "sync"
	
	"google.golang.org/grpc/codes"
    "google.golang.org/grpc"
	
    pb "ummsgo/pb"
)

type Server struct {
    pb.UnimplementedUserServiceServer
    mutex sync.Mutex
    users map[string]string
}

func (server *Server) RegisterUser(
	ctx context.Context,
	req *pb.RegisterUserRequest,
) (*pb.RegisterUserResponse, error) {
    server.mutex.Lock()
    defer server.mutex.Unlock()
	
    if _, exists := server.users[req.Username]; exists {
        return &pb.RegisterUserResponse{ResponseMessage: "User already exists"}, nil
    }
    server.users[req.Username] = req.Email
    return &pb.RegisterUserResponse{ResponseMessage: "User registered successfully"}, nil
}

func (server *Server) GetUser(
	ctx context.Context,
	req *pb.GetUserRequest,
) (*pb.GetUserResponse, error) {
    server.mutex.Lock()
    defer server.mutex.Unlock()
	
    email, exists := server.users[req.Username]
    if !exists {
        return nil, grpc.Errorf(codes.NotFound, "User not found")
    }
    return &pb.GetUserResponse{Username: req.Username, Email: email}, nil
}

func main() {
	port := ":6666"
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("error: failed to listen: %v", err)
    }
	
    server := grpc.NewServer()
    pb.RegisterUserServiceServer(server, &Server{users: make(map[string]string)})
    log.Println("info: server is running on port", port)
	
    if err := server.Serve(lis); err != nil {
        log.Fatalf("error: failed to serve: %v", err)
    }
}
