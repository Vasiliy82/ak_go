package repository

import (
	"lesson/internal/report/domain"
)

type User struct{}

func (u User) ByPeriod(from, to string) []domain.User {
	return []domain.User{
		domain.User{
			Name:      "Остап Бендер",
			Posts:     10,
			Shows:     3,
			Reactions: 1,
		},
		domain.User{
			Name:      "Шура Балаганов",
			Posts:     15,
			Shows:     6,
			Reactions: 5,
		},
		domain.User{
			Name:      "Ипполит Матвеевич",
			Posts:     7,
			Shows:     2,
			Reactions: 0,
		},
	}
}
