package main

import (
	pb "blog/proto"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := &BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
	log.Println(oid)
	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data}, //$set for update
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not Update DB",
		)
	}

	if res.MatchedCount == 0 { // there are no matching entry with given ID
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog with Id",
		)
	}
	return &emptypb.Empty{}, nil
}
