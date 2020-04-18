// vessel-service/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/fusidic/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	defaultHost = "localhost:27017"
)

func createDummyData(repo repository) {
	vessels := []*Vessel{
		{ID: "vessel01", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(context.Background(), v)
	}
}

func main() {
	srv := micro.NewService(
		micro.Name("vessel"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(context.Background(), uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	vesselCollection := client.Database("shippy").Collection("vessel")
	repository := &VesselRepository{
		vesselCollection,
	}

	createDummyData(repository)

	// 将VesselRepository中的数据、方法与handler.repository绑定
	pb.RegisterVesselServiceHandler(srv.Server(), &handler{repository})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
