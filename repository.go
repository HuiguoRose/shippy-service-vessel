package main

import (
	"context"
	pb "github.com/HuiguoRose/shippy-service-vessel/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Vessel struct {
	Id        string `json:"id"`
	Capacity  int32  `json:"capacity"`
	MaxWeight int32  `json:"max_weight"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
	OwnerId   string `json:"ownerId"`
}

type Vessels []*Vessel

type Specification struct {
	Capacity  int32 `json:"capacity"`
	MaxWeight int32 `json:"max_weight"`
}

func UnmarshalVessel(vessel *Vessel) *pb.Vessel {
	return &pb.Vessel{
		Id:        vessel.Id,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerId:   vessel.OwnerId,
	}
}

func MarshalVessel(vessel *pb.Vessel) *Vessel {
	return &Vessel{
		Id:        vessel.Id,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerId:   vessel.OwnerId,
	}
}
func UnmarshalVesselCollection(vessels Vessels) []*pb.Vessel {
	collection := make([]*pb.Vessel, 0)
	for _, vessel := range vessels {
		collection = append(collection, UnmarshalVessel(vessel))
	}
	return collection
}

func UnmarshalSpecification(specification *Specification) *pb.Specification {
	return &pb.Specification{
		Capacity:  specification.Capacity,
		MaxWeight: specification.MaxWeight,
	}
}

func MarshalSpecification(specification *pb.Specification) *Specification {
	return &Specification{
		Capacity:  specification.Capacity,
		MaxWeight: specification.MaxWeight,
	}
}

type Repository interface {
	FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error)
	Create(ctx context.Context, vessel *Vessel) error
	GetVessels(ctx context.Context) (Vessels, error)
}

// VesselRepository implementation
type VesselRepository struct {
	collection *mongo.Collection
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repository *VesselRepository) FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error) {
	filter := bson.D{{
		"capacity",
		bson.D{{
			"$gte",
			spec.Capacity,
		}},
	}, {
		"maxweight",
		bson.D{{
			"$gte",
			spec.MaxWeight,
		}},
	}}
	vessel := &Vessel{}
	x := repository.collection.FindOne(ctx, filter)
	log.Printf("==== %v\n", filter)
	log.Printf("==== %v\n", x)
	if err := x.Decode(vessel); err != nil {
		return nil, err
	}
	return vessel, nil
}

// Create a new vessel
func (repository *VesselRepository) Create(ctx context.Context, vessel *Vessel) error {
	_, err := repository.collection.InsertOne(ctx, vessel)
	return err
}

func (repository *VesselRepository) GetVessels(ctx context.Context) (Vessels, error) {
	cur, err := repository.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = cur.Close(ctx)
	}()
	var vessels Vessels
	for cur.Next(ctx) {
		var vessel *Vessel
		if err := cur.Decode(&vessel); err != nil {
			return nil, err
		}
		vessels = append(vessels, vessel)
	}
	return vessels, err
}
