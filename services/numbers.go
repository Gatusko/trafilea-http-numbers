package services

import (
	"github.com/Gatusko/trafilea-http-numbers/domain/model"
)

type NumberService interface {
	AddNumber(model.Numbers) (model.Value, error)
	GetNumber(int) (model.Value, error)
	GetAllNumbers() ([]model.Value, error)
	DeleteNumber(int) error
}
