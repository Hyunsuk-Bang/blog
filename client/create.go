package main

import (
	pb "blog/proto"
	"context"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	blog := &pb.Blog{
		AuthorId: "Hyunsuk Bang",
		Title:    "My First Article",
		Content:  "Content of the first Blog",
	}
	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("CreateBlog error! : %v", err)
	}
	return res.Id
}
