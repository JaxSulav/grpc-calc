package main

import (
	calc "calc/libs"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

const (
	address = "localhost:50051"
)

func main (){

	// Dial to the server address, the connection given by dial will be used to create a new calculator client
	conn, err := grpc.Dial(address , grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to server %v \n", err)
	}

 	cc := calc.NewCalculatorClient(conn)

	//sendUnary(cc)
	//serverStream(cc)
	clientStream(cc)
}

func sendUnary(cc calc.CalculatorClient) {
	a := int32(2)
	b := int32(3)

	req := &calc.SumRequest{
		A: a,
		B: b,
	}
	res, err := cc.SumService(context.Background(), req)

	if err != nil {
		log.Fatalf("Canot get response from the server: %v \n", err)
	}

	fmt.Printf("The Response from the server: %v \n", res.GetResult())

}

func serverStream(cc calc.CalculatorClient) {
	limit := uint32(100)
	req := &calc.PrimeRequest{
		Limit: limit,
	}

	stream, err := cc.PrimeService(context.Background(), req)

	if err != nil{
		log.Fatalf("Cannot receive streams from server: %v \n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF{
			fmt.Println("Calculating prime numbers finished")
			break
		}
		if err != nil{
			log.Fatalf("Cannot read value from stream: %v\n", err)
			break
		}

		fmt.Printf("Primes: %v \n", res.GetPrime())
	}
}


func clientStream(cc calc.CalculatorClient){
	nums := []float32{23.4, 22.3, 19.9, 11.11, 89.34, 78.44, 454.99}
	stream, err := cc.AverageService(context.Background())
	if err != nil {
		log.Fatalf("Canot create stream to send: %v", err)
	}

	for _, num := range nums{
		stream.Send(&calc.AverageRequest{
			Num: num,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Could not receive response from the server: %v", err)
	}

	log.Printf("The average from the server is: %v", res.GetAverage())
}

