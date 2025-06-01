package server_client

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/AshishNikam111000/gRPC/server_client/shop"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// in production use secure credentials
// always handle errors
func GrpcClient() {
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	c := shop.NewCoffeeShopClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	menuStream, err := c.GetMenu(ctx, &shop.MenuRequest{})
	if err != nil {
		log.Fatalf("Error calling function getMenu: %v", err)
	}

	done := make(chan bool)
	var items []*shop.Item
	go func() {
		for {
			resp, err := menuStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				log.Fatalf("can not recieve %v", err)
			}
			items = resp.Items
			log.Printf("Resp received: %v", items)
		}
	}()
	<-done

	receipt, _ := c.PlaceOrder(ctx, &shop.Order{Items: items})
	log.Printf("%v", receipt)

	status, _ := c.GetOrderStatus(ctx, receipt)
	log.Printf("%v", status)
}
