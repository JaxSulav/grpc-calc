package main

import (
	calc "calc/libs"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

const (
	port = "0.0.0.0:50051"
)

type server struct {
	calc.UnimplementedCalculatorServer

}

func calculateAndStreamPrime(limit uint32, stream calc.Calculator_PrimeServiceServer) int32 {
	var primes []uint32
	for i:=uint32(1); i<=limit; i++ {
		prime1 := i * 6 - 1
		prime2 := i * 6 + 1
		if prime1 > limit{
			break
		}
		primes = append(primes, prime1)
		primes = append(primes, prime2)
		log.Println("Sending stream of primes to client")
		time.Sleep(time.Millisecond * 500)
		stream.Send(&calc.PrimeResponse{Prime: primes})
	}
	return 0
}

func (*server) SumService(c context.Context, req *calc.SumRequest) (*calc.SumResponse, error){
	fmt.Println("Starting the Sum Service for calculator")
	a := req.GetA()
	b := req.GetB()
	fmt.Printf("The request from the client: %d, %d \n", a, b)

	result := a+b
	return &calc.SumResponse{
		Result: result,
	}, nil

}

func (*server) PrimeService(req *calc.PrimeRequest, stream calc.Calculator_PrimeServiceServer) error {
	limit := req.GetLimit()
	fmt.Printf("Recieved request to calculate prime number for limit: %v \n", limit)
	calculateAndStreamPrime(limit, stream)
	return nil
}

func (*server) AverageService(stream calc.Calculator_AverageServiceServer) error {
	average := float32(0)
	cnt := uint32(0)
	log.Println("Starting Average Calculator Service...")
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			stream.SendAndClose(&calc.AverageResponse{
				Average: average,
			})
			break
		}
		if err != nil {
			log.Printf("Could not receive from the streaming client: %v", err)
			break
		}
		reqNum := req.GetNum()
		log.Printf("Received stream for average: %v \n", reqNum)
		cnt++
		time.Sleep(time.Millisecond * 500)
		average = (average + reqNum) / float32(cnt)

	}
	log.Printf("Sending response average: %v", average)
	return nil
}

func main () {
	// Establish connection with the host
	lis, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("Could not establish connection at %v \n", port)
	}

	// instance of grpc server
	s := grpc.NewServer()
	calc.RegisterCalculatorServer(s, &server{})

	fmt.Printf("Listening at: %v \n", lis.Addr())
	// Serve to port through grpc Server
	if err := s.Serve(lis); err != nil{
		log.Fatalf("Could not server through grpc server: %v \n", err)
	}
}