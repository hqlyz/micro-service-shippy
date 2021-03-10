package main

import (
	"context"
	"encoding/json"
	"github.com/micro/micro/v3/service/client"
	"io/ioutil"
	"log"
	pb "micro-service-shippy/shippy-service-consignment/proto/consignment"
	"os"

)

const (
	address         = "localhost:5051"
	defaultFilename = "consignment.json"
)

func main() {
	// Set up a connection to the server.
	// conn, err := grpc.Dial(address, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("grpc dial err: %v\n", err)
	// }
	// defer conn.Close()

	// client := pb.NewShippingServiceClient(conn)
	cli := pb.NewShippingService("consignment", client.DefaultClient)

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("could not parse file: %v\n", err)
	}

	resp, err := cli.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("failed to create consignment: %v\n", err)
	}
	log.Printf("Created: %t\n", resp.Created)

	getAll, err := cli.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}

func parseFile(file string) (*pb.Consignment, error) {
	fileBuffer, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("failed to open file: %v\n", err)
		return nil, err
	}
	var consignment pb.Consignment
	err = json.Unmarshal(fileBuffer, &consignment)
	if err != nil {
		log.Printf("failed to unmarshal json file: %v\n", err)
		return nil, err
	}
	return &consignment, nil
}
