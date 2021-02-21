package main

import (
	"fmt"
	"time"

	grpc "justino.com/poc-client/grpc"
	rest "justino.com/poc-client/rest"
)

var (
	nAttempts = 50
)

func testGettingPayload(grpcClient grpc.ChatServiceClient, arg string, size grpc.Size, description string) {
	fmt.Printf("%v (%v):\n", description, arg)
	restStart := time.Now()
	for j := 0; j < nAttempts; j++ {
		rest.GetPayload(arg)
	}
	restDuration := time.Since(restStart)
	fmt.Printf("REST: %v \n", restDuration)

	grpcStart := time.Now()
	for i := 0; i < nAttempts; i++ {
		grpc.RunGetPayload(grpcClient, size)
	}
	grpcDuration := time.Since(grpcStart)
	fmt.Printf("GRPC: %v \n", grpcDuration)

	if restDuration > grpcDuration {
		fmt.Printf("GRPC was %v %% faster \n", 100*float64(restDuration.Milliseconds())/float64(grpcDuration.Milliseconds()))
	} else {
		fmt.Printf("Rest was %v %% faster \n", 100*float64(restDuration.Milliseconds())/float64(restDuration.Milliseconds()))
	}
	fmt.Println()
}

func main() {
	grpcClient := grpc.StartupClient()

	testGettingPayload(grpcClient, "1kb", grpc.Size_TINY, "testTiny")
	testGettingPayload(grpcClient, "500kb", grpc.Size_SMALL, "testSmall")
	testGettingPayload(grpcClient, "1mb", grpc.Size_MEDIUM, "testMedium")
	testGettingPayload(grpcClient, "10mb", grpc.Size_LARGE, "testLarge")
	testGettingPayload(grpcClient, "100mb", grpc.Size_HUGE, "testHuge")
}
