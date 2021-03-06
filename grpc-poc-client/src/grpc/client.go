package grpc

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

func RunGetPayload(client ChatServiceClient, size Size) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	desiredSize := PayloadSize{Size: size}
	_, err := client.GetPayload(ctx, &desiredSize)
	if err != nil {
		log.Fatalf("%v.runGetPayload(_) = _, %v: ", client, err)
	}
}

func runGetMessages(client ChatServiceClient) {
	log.Printf("Getting messages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := client.GetMessages(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("%v.GetMessages(_) = _, %v: ", client, err)
	}
	log.Println(response)
	log.Println("=========================")
}

func sendMessage(client ChatServiceClient, content string) {
	log.Printf("Sending message")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	msg := ChatMessage{
		Content: content,
	}
	response, err := client.SendMessage(ctx, &msg)
	if err != nil {
		log.Fatalf("%v.GetMessages(_) = _, %v: ", client, err)
	}
	log.Println(response)
	log.Println("=========================")
}

func listen(stream ChatService_ConsumeMessagesClient) {
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

func receiveNewMessages(client ChatServiceClient) {
	log.Printf("Receiving n messages")
	stream, _ := client.ConsumeMessages(context.Background(), &empty.Empty{})

	go listen(stream)
	time.Sleep(1 * time.Second)
	log.Println("=========================")

}

func runMakeRequests(client ChatServiceClient) {
	stream, _ := client.MakeRequests(context.Background())
	go listen(stream)

	request := ClientRequest{
		Action:  Action_SET_STATUS,
		Content: "Busy",
	}
	stream.Send(&request)
	time.Sleep(1 * time.Second)

	request = ClientRequest{
		Action:  Action_LOGOUT,
		Content: "Vlw flw",
	}
	stream.Send(&request)
	time.Sleep(1 * time.Second)
	log.Println("=========================")
}

func StartupClient() ChatServiceClient {
	flag.Parse()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithMaxMsgSize(512 * 1024 * 1024),
	}

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	client := NewChatServiceClient(conn)

	return client
}

func RunGRPCTests(client ChatServiceClient) {
	sendMessage(client, "xa")
	sendMessage(client, "blau")
	runGetMessages(client)
	receiveNewMessages(client)
	runMakeRequests(client)
}
