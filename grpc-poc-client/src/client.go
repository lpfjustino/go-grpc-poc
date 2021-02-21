package main

import (
	// grpc "justino.com/poc-client/grpc"
	rest "justino.com/poc-client/rest"
)

func main() {
	// grpc.StartupClient()
	rest.GetLargePayload()
}
