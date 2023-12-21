package usecase

import (
	"lesson/internal/report/domain"
)

type popularity struct{}

func (p popularity) calc(u domain.User) float32 {
	if u.Shows == 0 {
		return 0
	}
	return float32(u.Posts) * (float32(u.Reactions) / float32(u.Shows))
}
