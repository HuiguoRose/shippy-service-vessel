// shippy-service-vessel/main.go
package main

import (
	"context"
	"fmt"
	pb "github.com/HuiguoRose/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro"
	"log"
	"os"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	//vessels := []*pb.Vessel{
	//	&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	//}

	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	srv.Init()
	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("vessels")

	repository := &VesselRepository{consignmentCollection}
	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &handler{repository})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
