package main

import (
	"context"
	"log"
	"net"

	pb "github.com/gowithvikash/grpc_with_go/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Server struct {
	pb.BlogServiceServer
}

var (
	uri     = "mongodb://root:root@localhost:27017/"
	network = "tcp"
	address = "0.0.0.0:50051"
)

var collection *mongo.Collection

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("Blog_Database").Collection("Blog_Collection")
	// DbNames, err := client.ListDatabaseNames(context.Background(), bson.D{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("DbNames: %v\n", DbNames)
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}

func (s *Server) Create_New_Blog(ctx context.Context, data *pb.Blog) (*pb.BlogId, error) {
	log.Println("____  _Blog Function Was Invoked At Server  ____")
}

// func (s *Server) Read_Blog(ctx context.Context, id *pb.BlogId) (*pb.Blog, error) {
// 	log.Println("____  _Blog Function Was Invoked At Server  ____")
// }
// func (s *Server) Update_Blog(ctx context.Context, data *pb.Blog) (*emptypb.Empty, error) {
// 	log.Println("____  _Blog Function Was Invoked At Server  ____")
// }
// func (s *Server) Delete_Blog(ctx context.Context, id *pb.BlogId) (*emptypb.Empty, error) {
// 	log.Println("____  _Blog Function Was Invoked At Server  ____")
// }
// func (s *Server) List_All_Blogs(n *emptypb.Empty, stream pb.BlogService_List_All_BlogsServer) error {
// 	log.Println("____  _Blog Function Was Invoked At Server  ____")
// }
