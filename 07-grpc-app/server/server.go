package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/tkmagesh/cisco-advgo-apr-2023/07-grpc-app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (asi *appServiceImpl) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	for no := start; no <= end; no++ {
		if isPrime(no) {
			resp := &proto.PrimeResponse{
				PrimeNo: no,
			}
			log.Printf("Sending Prime no %d\n", no)
			serverStream.Send(resp)
			time.Sleep(500 * time.Millisecond)
		}
	}
	return nil
}

func (asi *appServiceImpl) CalculateAverage(serverStream proto.AppService_CalculateAverageServer) error {
	var sum, count int32
	for {
		avgReq, err := serverStream.Recv()
		if err == io.EOF {
			avg := sum / count
			resp := &proto.AverageResponse{
				Average: avg,
			}
			err := serverStream.SendAndClose(resp)
			if err != nil {
				log.Fatalln(err)
				return err
			}
			break
		}
		if err != nil {
			if code := status.Code(err); code == codes.Canceled {
				log.Println("Cancel request received....")
				return nil
			}
			return err
		}
		log.Printf("Received No : %d\n", avgReq.GetNo())
		sum += avgReq.GetNo()
		count++
	}

	return nil
}

/* Bidirectional Streaming */
func (asi *appServiceImpl) Greet(serverStream proto.AppService_GreetServer) error {
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		personName := *req.GetPerson()
		firstName := personName.GetFirstName()
		lastName := personName.GetLastName()
		fmt.Printf("Received req : %s %s\n", firstName, lastName)
		greetMsg := fmt.Sprintf("Hi %s %s, Have a nice day", firstName, lastName)
		fmt.Printf("Sending response : %s\n", greetMsg)
		time.Sleep(500 * time.Millisecond)
		res := &proto.GreetResponse{
			Message: greetMsg,
		}
		if err := serverStream.Send(res); err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}

func isPrime(no int32) bool {
	for i := int32(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
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
