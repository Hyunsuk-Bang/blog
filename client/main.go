package main

import (
	pb "blog/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("CLIENT: Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	blog := readBlog(c, id)
	log.Printf("%v\n", blog)

	updateBlog(c, id)
	blog = readBlog(c, id)
	log.Printf("%v\n", blog)

	listBlog(c)

	deleteBlog(c, id)

	listBlog(c)
}
