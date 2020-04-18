package main

import (
	"context"

	pb "github.com/fusidic/vessel-service/proto/vessel"
)

type handler struct {
	repository
}

// FindAvailable vessels
func (s *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	// 递归寻找合适的vessel
	vessel, err := s.repository.FindAvailable(ctx, MarshalSpecification(req))
	if err != nil {
		return err
	}

	// 将vessel添加到response信息中
	res.Vessel = UnmarshalVessel(vessel)
	return nil
}

// Create a new vessel
func (s *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := s.repository.Create(ctx, MarshalVessel(req)); err != nil {
		return err
	}
	res.Vessel = req
	return nil
}
