package main

import (
	"context"
	pb "github.com/HuiguoRose/shippy-service-vessel/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
			"$lte",
			spec.Capacity,
		}, {
			"$lte",
			spec.MaxWeight,
		}},
	}}
	vessel := &Vessel{}
	if err := repository.collection.FindOne(ctx, filter).Decode(vessel); err != nil {
		return nil, err
	}
	return vessel, nil
}

// Create a new vessel
func (repository *VesselRepository) Create(ctx context.Context, vessel *Vessel) error {
	_, err := repository.collection.InsertOne(ctx, vessel)
	return err
}
