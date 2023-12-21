package controller

import (
	"context"
	pb "lesson/pkg/report"
)

func (c *Controller) Healthz(ctx context.Context, req *pb.HealthzRequest) (*pb.HealthzResponse, error) {
	return &pb.HealthzResponse{}, nil
}
