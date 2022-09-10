package main

import (
	pb "blog/proto"
	"context"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	stream, err := c.ListBlog(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Printf("Cannot list blogs:%v\n", err)
	}

	for {
		blog, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error! %v", err)
		}
		log.Println(blog)
	}
}
