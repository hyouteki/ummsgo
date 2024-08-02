package main

import (
    "context"
    "log"
    "google.golang.org/grpc"

	pb "ummsgo/pb"
)

func registerUser(client pb.UserServiceClient, username, email string) {
    resp, err := client.RegisterUser(context.Background(),
		&pb.RegisterUserRequest{Username: username, Email: email})
    if err != nil {
        log.Fatalf("error: could not register user: %v", err)
    }
    log.Printf("info: register response: %s", resp.ResponseMessage)
}

func getUser(client pb.UserServiceClient, username string) {
    resp, err := client.GetUser(context.Background(),
		&pb.GetUserRequest{Username: username})
    if err != nil {
        log.Fatalf("error: could not get user: %v", err)
    }
    log.Printf("info: user: %s, email: %s", resp.Username, resp.Email)
}

func main() {
    conn, err := grpc.Dial("localhost:6666",
		grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("error: could not connect: %v", err)
    }
    defer conn.Close()
	
    client := pb.NewUserServiceClient(conn)

    registerUser(client, "nosferatu", "nosferatu@orava.castle")
    getUser(client, "nosferatu")
}
