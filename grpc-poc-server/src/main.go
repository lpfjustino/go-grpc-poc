package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	pb "justino.com/poc/grpc"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type chatServer struct {
	pb.UnimplementedChatServiceServer

	messages []string
}

func (s *chatServer) GetMessages(context.Context, *empty.Empty) (*pb.ChatMessage, error) {
	res := pb.ChatMessage{
		Content: "DEU BOM",
	}
	return &res, nil
}

func (s *chatServer) SendMessage(c context.Context, m *pb.ChatMessage) (*pb.ServerResponse, error) {
	res := pb.ServerResponse{
		Content: "Received",
	}
	log.Printf("Sending message %v", m.Content)
	s.messages = append(s.messages, m.Content)
	return &res, nil
}

func (s *chatServer) ConsumeMessages(e *empty.Empty, stream pb.ChatService_ConsumeMessagesServer) error {
	if len(s.messages) == 0 {
		log.Printf("No new messages")
		return nil
	}

	for _, message := range s.messages {
		response := pb.ServerResponse{
			Content: message,
		}
		if err := stream.Send(&response); err != nil {
			log.Printf("Error on receiving message: " + err.Error())
			return err
		}
	}
	s.messages = []string{}

	return nil
}

func (s *chatServer) MakeRequests(stream pb.ChatService_MakeRequestsServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf(in.Content, in.Action, in.GetAction(), in.GetContent())

		switch in.Action {
		case pb.Action_LOGOUT:
			response := wrapServerResponse("Going home, already?")
			stream.Send(&response)
		case pb.Action_SET_STATUS:
			response := wrapServerResponse("Status changed")
			stream.Send(&response)
		}
	}
}

func newServer() *chatServer {
	s := &chatServer{messages: make([]string, 0)}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterChatServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

func wrapServerResponse(message string) pb.ServerResponse {
	return pb.ServerResponse{
		Content: message,
	}
}
