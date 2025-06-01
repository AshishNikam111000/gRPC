package server_client

import (
	"context"
	"log"
	"net"

	"github.com/AshishNikam111000/gRPC/server_client/shop"
	"google.golang.org/grpc"
)

type coffeshop struct {
	shop.UnimplementedCoffeeShopServer
}

func (s *coffeshop) GetMenu(menuReq *shop.MenuRequest, getMenuServer shop.CoffeeShop_GetMenuServer) error {
	items := []*shop.Item{
		{
			Id:   "1",
			Name: "Black Coffee",
		},
		{
			Id:   "2",
			Name: "Americano",
		},
		{
			Id:   "3",
			Name: "Latte",
		},
	}
	for i := range items {
		getMenuServer.Send(&shop.Menu{Items: items[0 : i+1]})
	}
	return nil
}
func (s *coffeshop) PlaceOrder(ctx context.Context, order *shop.Order) (*shop.Receipt, error) {
	return &shop.Receipt{Id: "123"}, nil
}
func (s *coffeshop) GetOrderStatus(ctx context.Context, receipt *shop.Receipt) (*shop.OrderStatus, error) {
	return &shop.OrderStatus{
		OrderId: receipt.Id,
		Status:  "IN PROGRESS",
	}, nil
}

func GrpcServer() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Listener started on tcp localhost:9000")

	grpcServe := grpc.NewServer()
	shop.RegisterCoffeeShopServer(grpcServe, &coffeshop{})
	if err := grpcServe.Serve(lis); err != nil {
		log.Fatalf("Failed to serve the gRPC server: %v", err)
	}
	log.Println("gRPC server started")
}
