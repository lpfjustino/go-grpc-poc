package main

import (
	"fmt"
	"log"
	"time"

	grpc "justino.com/poc-client/grpc"
	rest "justino.com/poc-client/rest"
)

var (
	nAttempts = 10
)

func testTiny(grpcClient grpc.ChatServiceClient) {
	log.Println("testTiny")
	start := time.Now()
	for j := 0; j < nAttempts; j++ {
		rest.GetPayload("1kb")
	}
	duration := time.Since(start)
	fmt.Println(duration)

	start = time.Now()
	for i := 0; i < nAttempts; i++ {
		grpc.RunGetPayload(grpcClient, grpc.Size_TINY)
	}
	duration = time.Since(start)
	fmt.Println(duration)
}

func testSmall(grpcClient grpc.ChatServiceClient) {
	log.Println("testSmall")
	start := time.Now()
	for j := 0; j < nAttempts; j++ {
		rest.GetPayload("500kb")
	}
	duration := time.Since(start)
	fmt.Println(duration)

	start = time.Now()
	for i := 0; i < nAttempts; i++ {
		grpc.RunGetPayload(grpcClient, grpc.Size_SMALL)
	}
	duration = time.Since(start)
	fmt.Println(duration)
}

func testMedium(grpcClient grpc.ChatServiceClient) {
	log.Println("testMedium")
	start := time.Now()
	for j := 0; j < nAttempts; j++ {
		rest.GetPayload("1mb")
	}
	duration := time.Since(start)
	fmt.Println(duration)

	start = time.Now()
	for i := 0; i < nAttempts; i++ {
		grpc.RunGetPayload(grpcClient, grpc.Size_MEDIUM)
	}
	duration = time.Since(start)
	fmt.Println(duration)
}

func testLarge(grpcClient grpc.ChatServiceClient) {
	log.Println("testLarge")
	start := time.Now()
	for j := 0; j < nAttempts; j++ {
		rest.GetPayload("10mb")
	}
	duration := time.Since(start)
	fmt.Println(duration)

	start = time.Now()
	for i := 0; i < nAttempts; i++ {
		grpc.RunGetPayload(grpcClient, grpc.Size_LARGE)
	}
	duration = time.Since(start)
	fmt.Println(duration)
}

func testHuge(grpcClient grpc.ChatServiceClient) {
	log.Println("testHuge")
	start := time.Now()
	for j := 0; j < nAttempts; j++ {
		rest.GetPayload("100mb")
	}
	duration := time.Since(start)
	fmt.Println(duration)

	start = time.Now()
	for i := 0; i < nAttempts; i++ {
		grpc.RunGetPayload(grpcClient, grpc.Size_HUGE)
	}
	duration = time.Since(start)
	fmt.Println(duration)
}

func main() {
	grpcClient := grpc.StartupClient()
	testTiny(grpcClient)
	testSmall(grpcClient)
	testMedium(grpcClient)
	testLarge(grpcClient)
	testHuge(grpcClient)
}
