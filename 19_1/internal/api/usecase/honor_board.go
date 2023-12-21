package usecase

import (
	"context"
	"lesson/internal/api/domain"
	"lesson/internal/api/repository"
	"sort"
)

type HonorBoard struct {
	source repository.Popularity
}

func NewHonorBoard(s repository.Popularity) HonorBoard {
	return HonorBoard{
		source: s,
	}
}
func (hb HonorBoard) Users(ctx context.Context, from, to string) []domain.User {
	p, err := hb.source.ByPeriod(ctx, from, to)
	if err != nil {
		return []domain.User{}
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].Value > p[j].Value
	})
	result := make([]domain.User, 0, len(p))
	for _, v := range p {
		if v.Value == 0 {
			continue
		}
		result = append(result, domain.User{
			Name: v.User.Name,
		})
	}
	return result
}
