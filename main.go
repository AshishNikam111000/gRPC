package main

import (
	"sync"

	"github.com/AshishNikam111000/gRPC/server_client"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		defer wg.Done()
		server_client.GrpcServer()
	}()
	go func() {
		defer wg.Done()
		server_client.GrpcClient()
	}()
}
