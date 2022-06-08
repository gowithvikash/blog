package main

import (
	pb "github.com/gowithvikash/grpc_with_go/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Author_Id string             `bson:"author_id"`
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
}

func dataToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.Author_Id,
		Title:    data.Title,
		Content:  data.Content,
	}

}
