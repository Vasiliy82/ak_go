package usecase

import (
	"lesson/internal/report/domain"
	"lesson/internal/report/repository"
)

type User struct {
	source     repository.User
	popularity popularity
}

func (u User) Info(from, to string) []domain.UserInfo {
	users := u.source.ByPeriod(from, to)
	result := make([]domain.UserInfo, 0, len(users))
	for _, v := range users {
		result = append(result, domain.UserInfo{
			Name:       v.Name,
			Popularity: u.popularity.calc(v),
		})
	}
	return result
}
