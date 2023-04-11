package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/tkmagesh/cisco-advgo-apr-2023/07-grpc-app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}
	appServiceClient := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()

	// doRequestResponse(ctx, appServiceClient)
	doServerStreaming(ctx, appServiceClient)
}

func doRequestResponse(ctx context.Context, appServiceClient proto.AppServiceClient) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	res, err := appServiceClient.Add(timeoutCtx, addRequest)
	if err != nil {
		if code := status.Code(err); code == codes.DeadlineExceeded {
			fmt.Println("Timeout occured")
			return
		}
		log.Fatalln(err)
	}
	fmt.Println("Result =", res.GetResult())
}

func doServerStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) {
	primeReq := &proto.PrimeRequest{
		Start: 3,
		End:   100,
	}
	clientStream, err := appServiceClient.GeneratePrimes(ctx, primeReq)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		resp, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("All the prime nos received")
			return
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Prime No :", resp.GetPrimeNo())
	}

}
