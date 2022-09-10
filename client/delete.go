package main

import (
	pb "blog/proto"
	"context"
	"log"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error while deleting: %v", err)
	}

	log.Println("Blog was deleted!")
}
