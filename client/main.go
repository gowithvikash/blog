package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/gowithvikash/grpc_with_go/blog/proto"
	"google.golang.org/grpc"
)

var (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)
	Create_New_Blog(c)
	Read_Blog(c)
	Update_Blog(c)
	Delete_Blog(c)
	List_All_Blogs(c)

}

func Create_New_Blog(c pb.BlogServiceClient) {
	log.Println("____  Create_New_Blog Function Was Invoked At Client  ____")
	var data = &pb.Blog{
		AuthorId: "Vikash Parashar",
		Title:    "Far From Home",
		Content:  "Try To Avoid Such Type Of Content",
	}
	bid, err := c.Create_New_Blog(context.Background(), data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bid)
}
func Read_Blog(c pb.BlogServiceClient) {
	log.Println("____  Read_Blog Function Was Invoked At Client  ____")
}
func Update_Blog(c pb.BlogServiceClient) {
	log.Println("____  Update_Blog Function Was Invoked At Client  ____")
}
func Delete_Blog(c pb.BlogServiceClient) {
	log.Println("____  Delete_Blog Function Was Invoked At Client  ____")
}
func List_All_Blogs(c pb.BlogServiceClient) {
	log.Println("____  List_All_Blogs Function Was Invoked At Client  ____")
}
