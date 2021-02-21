package main

import (
	"flag"

	sv "justino.com/poc/grpc"
	rest "justino.com/poc/rest"
)

func main() {
	flag.Parse()
	go sv.ServerStartup()
	rest.StartupRestServer()
}
