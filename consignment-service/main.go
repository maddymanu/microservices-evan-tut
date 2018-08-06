package main

import (
	"os"
	vesselProto "github.com/maddymanu/microservices-evan-tut/vessel-service/proto/vessel"
	userService "github.com/maddymanu/microservices-evan-tut/user-service/proto/user"
	pb "github.com/maddymanu/microservices-evan-tut/consignment-service/proto/consignment"
	"github.com/micro/go-micro"
	"fmt"
	"github.com/micro/go-micro/server"
	"context"
	"github.com/micro/go-micro/metadata"
	"log"
	"errors"
	"github.com/micro/go-micro/client"
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
	}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
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

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc  {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)

		if !ok {
			log.Println("bad things happened while parsing metadata")
			return errors.New("no auth metadata found, bad request")
		}

		log.Println(meta)

		token := meta["Token"]
		log.Println("Found token: ", token)

		authClient := userService.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &userService.Token{
			Token:token,
		})

		if err != nil {
			return err
		}

		err = fn(ctx, req, resp)
		return err
	}
}