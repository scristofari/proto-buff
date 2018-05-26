package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/scristofari/proto/poll"
	"google.golang.org/grpc"
)

type pollServer struct{}

func (ps *pollServer) Get(context.Context, *poll.PollRequest) (*poll.Poll, error) {
	i := &timestamp.Timestamp{}
	i.Seconds = time.Now().Unix()
	p := &poll.Poll{
		Id:           1,
		Title:        "test",
		LastModified: i,
	}
	return p, nil
}

func (ps *pollServer) Ticker(e *empty.Empty, stream poll.PollServer_TickerServer) error {
	for i := 0; i < 5; i++ {
		stream.Send(&poll.TickerResponse{Tick: fmt.Sprintf("%d", i)})
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	s := grpc.NewServer()
	poll.RegisterPollServerServer(s, &pollServer{})
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("could not listen to :8888: %v", err)
	}
	log.Fatal(s.Serve(l))
}
