package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/serdarkalayci/go-grpc/product-api/proto"
	"google.golang.org/grpc"
)

// printProducts demonstrates simple binary call from client
func printProducts(client pb.ProductServiceClient) {
	var emptyReq empty.Empty
	emptyReq = empty.Empty{}
	curList, err := client.GetProducts(context.Background(), &emptyReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nProducts")
	fmt.Println("-------------")
	for _, cur := range curList.Products {
		fmt.Printf("%-50s%-10s\n", cur.Name, cur.Description)
	}
}

// selectProduct demonstrates simple binary call from client
func selectProduct(client pb.ProductServiceClient) {
	prodReq := &pb.ProductQuery{
		Id: 1,
	}
	item, err := client.GetProduct(context.Background(), prodReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nProduct")
	fmt.Println("-------------")
	fmt.Printf("%-50s%-10s\n", item.Name, item.Description)
}

func main() {
	serverAddr := net.JoinHostPort("localhost", "9001")

	// setup insecure connection
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewProductServiceClient(conn)

	printProducts(client)
	selectProduct(client)
}
