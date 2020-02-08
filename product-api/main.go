package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/serdarkalayci/go-grpc/product-api/proto"
)

type ProductService struct {
	data []*pb.Product
}

func newProductService(data []*pb.Product) *ProductService {
	return &ProductService{data: data}
}

// GetProducts searches (by Code or Number) and return ProductList
func (c *ProductService) GetProducts(
	ctx context.Context,
	e *empty.Empty,
	//req *pb.ProductRequest,
) (*pb.ProductList, error) {

	// var items []*pb.Product
	// for _, cur := range c.data {
	// 	if cur.GetNumber() == req.GetNumber() || cur.GetCode() == req.GetCode() {
	// 		items = append(items, cur)
	// 	}
	// }

	// return &pb.ProductList{Items: items}, nil
	return &pb.ProductList{Products: c.data}, nil
}

func main() {

	lstnr, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal("failed to start server:", err)
	}

	// setup and register Product service
	productService := newProductService(productList)
	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, productService)

	// start service's server
	log.Println("starting Product rpc service on", 9000)
	if err := grpcServer.Serve(lstnr); err != nil {
		log.Fatal(err)
	}
}

var productList = []*pb.Product{
	&pb.Product{
		Id:          1,
		Name:        "Latte",
		Description: "Froth milky coffee",
		Price:       1,
	},
	&pb.Product{
		Id:          2,
		Name:        "Espresso",
		Description: "Black coffee",
		Price:       1,
	},
}
