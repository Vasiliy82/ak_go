package controller

import (
	"context"
	"lesson/internal/report/usecase"
	pb "lesson/pkg/report"
)

func (c *Controller) User(ctx context.Context, req *pb.Dates) (*pb.Users,
	error) {
	uc := usecase.User{}
	data := uc.Info(req.Start, req.End)
	res := pb.Users{
		User: make([]*pb.User, 0, len(data)),
	}
	for _, v := range data {
		u := pb.User{
			Name:       v.Name,
			Popularity: v.Popularity,
		}
		res.User = append(res.User, &u)
	}
	return &res, nil
}
