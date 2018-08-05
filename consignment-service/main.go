package main

import (
	pb "github.com/maddymanu/microservices-evan-tut/consignment-service/proto/consignment"
	vesselProto "github.com/maddymanu/microservices-evan-tut/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"github.com/labstack/gommon/log"
	"fmt"
	"os"
)

const (
	defaultHost = "localhost:27017"
)
// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.

func main() {

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	// Mgo creates a 'master' session, we need to end that session
	// before the main function closes.
	defer session.Close()

	if err != nil {

		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)

		// Create a new service. Optionally include some options here.
		srv := micro.NewService(

			// This name must match the package name given in your protobuf definition
			micro.Name("go.micro.srv.consignment"),
			micro.Version("latest"),
		)

		vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

		// Init will parse the command line flags.
		srv.Init()

		pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

		// Run the server
		if err := srv.Run(); err != nil {
			fmt.Println(err)
		}
	}

}