package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/scristofari/proto/poll"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v\n", err)
		os.Exit(1)
	}
	client := poll.NewPollServerClient(conn)

	p, _ := client.Get(context.Background(), &poll.PollRequest{Id: 1})
	fmt.Println(p)

	stream, _ := client.Ticker(context.Background(), &empty.Empty{})
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v", err)
		}
		fmt.Println(data)
	}
}
