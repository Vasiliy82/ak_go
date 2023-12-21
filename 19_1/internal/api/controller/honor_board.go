package controller

import (
	"encoding/json"
	"lesson/internal/api/domain"
	"net/http"
)

func (c Controller) HonorBoard(res http.ResponseWriter, req *http.Request) {
	dates := Dates{}
	if err := json.NewDecoder(req.Body).Decode(&dates); err != nil {
		c.log.Println("ошибка парсинга:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	users := c.honorBoard.Users(req.Context(), dates.From, dates.To)
	result := domain.HonorBoard{
		Style: "Новинки в IT",
		Users: users,
	}
	err := json.NewEncoder(res).Encode(result)
	if err != nil {
		c.log.Println("ошибка парсинга:", err)
	}
}
