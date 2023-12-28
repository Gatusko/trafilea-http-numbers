package services

import (
	"github.com/Gatusko/trafilea-http-numbers/domain/model"
	"github.com/Gatusko/trafilea-http-numbers/domain/repositories"
	"log"
)

type numberService struct {
	numberRepository repositories.NumbersRepository
}

func NewNumberSerivce(numberRepository repositories.NumbersRepository) *numberService {
	return &numberService{
		numberRepository: numberRepository,
	}
}

func (ns *numberService) AddNumber(number model.Numbers) (model.Value, error) {
	err := number.Validate()
	if err != nil {
		return model.Value{}, err
	}
	value := model.NewValue(number.Value)
	val, err := ns.numberRepository.Add(number, value)
	if err != nil {
		log.Printf("Error adding new Number  to the memory :%v", err)
		return model.Value{}, err
	}
	return val, nil
}

func (ns *numberService) GetNumber(numberToCheck int) (model.Value, error) {
	val, err := ns.numberRepository.Get(numberToCheck)
	if err != nil {
		return model.Value{}, err
	}
	return val, nil
}

func (ns *numberService) GetAllNumbers() ([]model.Value, error) {
	vals, err := ns.numberRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return vals, nil
}

func (ns *numberService) DeleteNumber(number int) error {
	return ns.numberRepository.Delete(number)
}
