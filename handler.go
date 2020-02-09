package main

import (
	"context"
	pb "github.com/HuiguoRose/shippy-service-vessel/proto/vessel"
	"log"
)

// Our grpc service handler
type handler struct {
	repo Repository
}

func (s *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	log.Printf("===req %s", req)
	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(ctx, MarshalSpecification(req))
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = UnmarshalVessel(vessel)
	log.Printf("===res %s", res)
	return nil
}

func (s *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	log.Printf("===req %s", req)
	if err := s.repo.Create(ctx, MarshalVessel(req)); err != nil {
		return err
	}
	res.Created = true
	res.Vessel = req
	log.Printf("===res %s", res)
	return nil
}

func (s *handler) GetVessels(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	log.Printf("===req %s", req)
	vessels, err := s.repo.GetVessels(ctx)
	if err != nil {
		return err
	}
	res.Vessels = UnmarshalVesselCollection(vessels)
	log.Printf("===res %s", res)
	return nil
}
