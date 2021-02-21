/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a simple gRPC client that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It interacts with the route guide service whose definition can be found in routeguide/route_guide.proto.
package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	pb "justino.com/poc-client/grpc"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func runGetMessages(client pb.ChatServiceClient) {
	log.Printf("Getting messages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := client.GetMessages(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("%v.GetMessages(_) = _, %v: ", client, err)
	}
	log.Println(response)
}

func sendMessage(client pb.ChatServiceClient, content string) {
	log.Printf("Sending message")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	xesque := pb.ChatMessage{
		Content: content,
	}
	response, err := client.SendMessage(ctx, &xesque)
	if err != nil {
		log.Fatalf("%v.GetMessages(_) = _, %v: ", client, err)
	}
	log.Println(response)
}

func listen(stream pb.ChatService_ConsumeMessagesClient) {
	waitc := make(chan struct{})
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			close(waitc)
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive a note : %v", err)
		}

		log.Printf("Recebi %v", in.Content)
	}
}

func receiveNewMessages(client pb.ChatServiceClient) {
	log.Printf("Receiving n messages")
	stream, _ := client.ConsumeMessages(context.Background(), &empty.Empty{})

	go listen(stream)
}

func runMakeRequests(client pb.ChatServiceClient) {
	stream, _ := client.MakeRequests(context.Background())
	go listen(stream)

	request := pb.ClientRequest{
		Action:  pb.Action_SET_STATUS,
		Content: "Busy",
	}
	stream.Send(&request)
	time.Sleep(1 * time.Second)

	request = pb.ClientRequest{
		Action:  pb.Action_LOGOUT,
		Content: "Vlw flw",
	}
	stream.Send(&request)
	time.Sleep(1 * time.Second)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewChatServiceClient(conn)

	// runGetMessages(client)
	// sendMessage(client)
	// runMakeRequests(client)

	sendMessage(client, "xa")
	sendMessage(client, "blau")
	receiveNewMessages(client)
	receiveNewMessages(client)

	time.Sleep(2 * time.Second)

}
