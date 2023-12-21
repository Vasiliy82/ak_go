package controller

import (
	"lesson/internal/api/usecase"
)

type Logger interface {
	Println(v ...any)
}
type Controller struct {
	log        Logger
	honorBoard usecase.HonorBoard
}

func NewController(hb usecase.HonorBoard, l Logger) Controller {
	return Controller{
		log:        l,
		honorBoard: hb,
	}
}
