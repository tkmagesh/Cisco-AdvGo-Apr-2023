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
	// doServerStreaming(ctx, appServiceClient)
	doClientStreaming(ctx, appServiceClient)
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

func doClientStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) {
	// nos := []int32{3, 1, 4, 2, 5}
	fmt.Println("Hit ENTER to cancel....")
	clientStream, err := appServiceClient.CalculateAverage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for no := int32(1); no <= 100; no++ {
		time.Sleep(500 * time.Millisecond)
		log.Printf("Sending No : %d\n", no)
		req := &proto.AverageRequest{
			No: no,
		}
		err := clientStream.Send(req)
		if err != nil {
			log.Fatalln(err)
		}
	}
	res, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Average : ", res.GetAverage())
}
