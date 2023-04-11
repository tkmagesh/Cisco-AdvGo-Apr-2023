package main

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/tkmagesh/cisco-advgo-apr-2023/07-grpc-app/proto"
	"google.golang.org/grpc"
)

type appServiceImpl struct {
	proto.UnimplementedAppServiceServer
}

func (asi *appServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	log.Printf("Add request received. x = %d and y = %d\n", x, y)
	// time.Sleep(5 * time.Second)
	select {
	case <-ctx.Done():
		log.Println("timeout occurred")
		return nil, errors.New("timeout occurred")
	default:
		result := x + y
		res := &proto.AddResponse{
			Result: result,
		}
		log.Println("Sending add response")
		return res, nil
	}

}

func main() {
	asi := &appServiceImpl{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
