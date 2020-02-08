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

// printUSD demonstrates simple binary call from client
func printUSD(client pb.ProductServiceClient) {
	var emptyReq empty.Empty
	emptyReq = empty.Empty{}
	curList, err := client.GetProducts(context.Background(), &emptyReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nUSD Countries")
	fmt.Println("-------------")
	for _, cur := range curList.Products {
		fmt.Printf("%-50s%-10s\n", cur.Name, cur.Description)
	}
}

func main() {
	serverAddr := net.JoinHostPort("localhost", "9001")

	// setup insecure connection
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewProductServiceClient(conn)

	printUSD(client)
}
