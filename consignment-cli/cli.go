package main


import (
	"encoding/json"
	"io/ioutil"
	pb "github.com/maddymanu/microservices-evan-tut/consignment-service/proto/consignment"
	"log"
	"os"
	"context"
	"github.com/micro/go-micro/cmd"
	microclient "github.com/micro/go-micro/client"

	"github.com/micro/go-micro/metadata"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main()  {

	cmd.Init()

	// Create new greeter client
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	// Contact the server and print out its response.
	file := defaultFilename
	var token string
	log.Println(os.Args)

	log.Println("len is " , len(os.Args))
	if len(os.Args) > 1 {
		file = os.Args[1]
		token = os.Args[3]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	// Create a new context which contains our given token.
	// This same context will be passed into both the calls we make
	// to our consignment-service.

	log.Println("Adding meta token" , token)

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	// First call using our tokenised context
	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	// Second call
	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}