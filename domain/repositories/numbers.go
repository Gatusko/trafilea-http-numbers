package repositories

import "github.com/Gatusko/trafilea-http-numbers/domain/model"

type NumbersRepository interface {
	Add(model.Numbers, model.Value) (model.Value, error)
	Get(int) (model.Value, error)
	GetAll() ([]model.Value, error)
	Delete(int) error
}
