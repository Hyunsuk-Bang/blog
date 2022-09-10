package main

import (
	pb "blog/proto"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlog(in *emptypb.Empty, stream pb.BlogService_ListBlogServer) error {
	log.Println("ListBlog was invoked")

	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			"Unknown internal error",
		)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				"Error while decoding data from MongoDB",
			)
		}

		stream.Send(documentToBlog(data))
	}
	if err = cur.Err(); err != nil {
		return status.Error(
			codes.Internal,
			"Unkown internal Error",
		)
	}
	return nil
}
