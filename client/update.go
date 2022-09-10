package main

import (
	pb "blog/proto"
	"context"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Hyunsuk Bang",
		Title:    "Updated first title",
		Content:  "Content of the first blog! UPDATED!!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error while updating:%v", err)
	}

}
