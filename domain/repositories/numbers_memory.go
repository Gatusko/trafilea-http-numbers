package repositories

import (
	"fmt"
	"github.com/Gatusko/trafilea-http-numbers/domain/model"
	"sync"
)

type NumbersMemory struct {
	memory map[int]model.Value
	mux    sync.RWMutex
}

func NewNumberMemoryRepository() *NumbersMemory {
	return &NumbersMemory{
		memory: make(map[int]model.Value),
	}
}

func (nm *NumbersMemory) Add(number model.Numbers, value model.Value) (model.Value, error) {
	nm.mux.Lock()
	nm.memory[number.Value] = value
	nm.mux.Unlock()
	return value, nil
}

func (nm *NumbersMemory) Get(number int) (model.Value, error) {
	nm.mux.Lock()
	num, ok := nm.memory[number]
	nm.mux.Unlock()
	if !ok {
		return model.Value{}, fmt.Errorf("Value not found: %v", number)
	}
	return num, nil
}
func (nm *NumbersMemory) GetAll() ([]model.Value, error) {
	numbers := []model.Value{}
	for _, v := range nm.memory {
		numbers = append(numbers, v)
	}
	return numbers, nil
}

func (nm *NumbersMemory) Delete(num int) error {
	_, err := nm.Get(num)
	if err != nil {
		return fmt.Errorf("Value not found: %v", num)
	}
	delete(nm.memory, num)
	return nil
}
