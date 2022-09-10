package main

import (
	pb "blog/proto"
	"context"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"
var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

func main() {
	//Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:password@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("blogdb").Collection("blog")

	// Open tcp connection
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("SERVER: Error listening from %s\n ERROR: %v", addr, err)
	}
	log.Printf("Listening from %s", addr)

	opts := []grpc.ServerOption{}
	tls := true
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterBlogServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("SERVER: Failed to serve: %v", err)
	}
}
