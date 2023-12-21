package repository

import (
	"context"
	"lesson/internal/api/domain"
	pb "lesson/pkg/report"
	"log"
)

type Logger interface {
	Println(v ...any)
}

type Popularity struct {
	log    Logger
	source pb.ReportClient
}

func NewPopularity(l Logger, s pb.ReportClient) Popularity {
	return Popularity{
		log:    l,
		source: s,
	}
}
func (p Popularity) ByPeriod(ctx context.Context, from, to string) ([]domain.Popularity, error) {
	req := pb.Dates{
		Start: from,
		End:   to,
	}
	resp, err := p.source.User(ctx, &req)
	if err != nil {
		log.Println(err)
		return []domain.Popularity{}, err
	}
	result := make([]domain.Popularity, 0, len(resp.User))
	for _, v := range resp.User {
		p := domain.Popularity{
			User: domain.User{
				Name: v.Name,
			},
			Value: int(v.Popularity),
		}
		result = append(result, p)
	}
	return result, err
}
