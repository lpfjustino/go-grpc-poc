package main

import (
	"fmt"
	"time"

	grpc "justino.com/poc-client/grpc"
	rest "justino.com/poc-client/rest"
)

func main() {
	grpcClient := grpc.StartupClient()

	start := time.Now()
	for j := 0; j < 1000; j++ {
		rest.GetLargePayload()
	}
	duration := time.Since(start)
	fmt.Println(duration)

	start = time.Now()
	for i := 0; i < 1000; i++ {
		grpc.RunGetLargePayload(grpcClient)
	}
	duration = time.Since(start)
	fmt.Println(duration)

}
