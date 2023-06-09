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
	// doClientStreaming(ctx, appServiceClient)
	doBidirectionalStreaming(ctx, appServiceClient)
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
	stopCh := make(chan struct{})
	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		fmt.Scanln()
		close(stopCh)
	}()

	clientStream, err := appServiceClient.CalculateAverage(cancelCtx)
	if err != nil {
		log.Fatalln(err)
	}
LOOP:
	for no := int32(1); no <= 100; no++ {
		select {
		case <-stopCh:
			cancel()
			break LOOP
		default:
			time.Sleep(100 * time.Millisecond)
			log.Printf("Sending No : %d\n", no)
			req := &proto.AverageRequest{
				No: no,
			}
			err := clientStream.Send(req)
			if err == io.EOF {
				fmt.Println("EOF received")
				break LOOP
			}
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	select {
	case <-stopCh: //<-cancelCtx.Done():
		fmt.Println("Operation cancelled...")
	default:
		res, err := clientStream.CloseAndRecv()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Average : ", res.GetAverage())
	}

}

func doBidirectionalStreaming(ctx context.Context, client proto.AppServiceClient) {
	personNames := []*proto.PersonName{
		&proto.PersonName{FirstName: "Magesh", LastName: "Kuppan"},
		&proto.PersonName{FirstName: "Suresh", LastName: "Kannan"},
		&proto.PersonName{FirstName: "Ganesh", LastName: "Kumar"},
		&proto.PersonName{FirstName: "Rajesh", LastName: "Pandit"},
		&proto.PersonName{FirstName: "Ramesh", LastName: "Jayaraman"},
	}
	clientStream, err := client.Greet(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	done := make(chan struct{})
	go func() {
		for {
			res, err := clientStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(res.GetMessage())
		}
		// done <- struct{}{}
		close(done)
	}()
	for _, personName := range personNames {
		fmt.Printf("Sending name : %s %s\n", personName.FirstName, personName.LastName)
		req := &proto.GreetRequest{
			Person: personName,
		}
		if err := clientStream.Send(req); err != nil {
			log.Fatalln(err)
		}
	}
	clientStream.CloseSend()
	<-done
	fmt.Println("Done!")
}
