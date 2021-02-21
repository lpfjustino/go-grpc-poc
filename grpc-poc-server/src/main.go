package main

import (
	"flag"

	sv "justino.com/poc/grpc"
)

func main() {
	flag.Parse()
	sv.ServerStartup()
}
