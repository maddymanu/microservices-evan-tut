package main


import (
	pb "github.com/maddymanu/microservices-evan-tut/vessel-service/proto/vessel"
	"os"
	"github.com/labstack/gommon/log"
	"github.com/micro/go-micro"
)
const (
	defaultHost = "localhost:27017"
)

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}


func main() {
	mongoHost := os.Getenv("DB_HOST")

	if mongoHost == "" {
		mongoHost = defaultHost
	}

	mongoSession, err := CreateSession(mongoHost)
	defer mongoSession.Close()

	if err != nil {
		log.Fatal("Error connetincting to mongo db")
	}

	repo := &VesselRepository{
		session:mongoSession.Copy(),
	}

	createDummyData(repo)

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{mongoSession})

	if err := srv.Run(); err != nil {
		log.Fatal("Fatal micro run")
	}
}