package services

import (
	"fmt"
	"github.com/Gatusko/trafilea-http-numbers/domain/model"
	"github.com/Gatusko/trafilea-http-numbers/domain/repositories"
	"testing"
)

func TestNumberService_AddNumber(t *testing.T) {
	tests := []struct {
		input         model.Numbers
		result        model.Value
		expectedError error
	}{
		{
			model.Numbers{5},
			model.Value{"Type 2"},
			nil,
		},
		{
			model.Numbers{3},
			model.Value{"Type 1"},
			nil,
		},
		{
			model.Numbers{15},
			model.Value{"Type 3"},
			nil,
		},
		{
			model.Numbers{1},
			model.Value{1},
			nil,
		},
	}
	numRepository := repositories.NewNumberMemoryRepository()
	numService := NewNumberSerivce(numRepository)
	for _, test := range tests {
		number, err := numService.AddNumber(test.input)
		if test.expectedError != nil && err != nil {
			if err.Error() != test.expectedError.Error() {
				t.Fatalf("Expected error: %s", test.expectedError)
			}
			continue
		}
		if number != test.result {
			t.Fatalf("Number is not matching correctly: %v != %v", number, test.result)
		}
	}
}

func TestNumberService_GetNumber(t *testing.T) {
	tests := []struct {
		input         model.Numbers
		result        model.Value
		search        int
		expectedError error
	}{
		{
			model.Numbers{1},
			model.Value{1},
			1,
			nil,
		},
		{
			model.Numbers{3},
			model.Value{"Type 1"},
			3,
			nil,
		},
		{
			model.Numbers{5},
			model.Value{"Type 2"},
			5,
			nil,
		},
		{
			model.Numbers{2},
			model.Value{0},
			555,
			fmt.Errorf("Value not found: %v", 555),
		},
	}
	numRepository := repositories.NewNumberMemoryRepository()
	numService := NewNumberSerivce(numRepository)
	for _, test := range tests {
		numService.AddNumber(test.input)
		number, err := numService.GetNumber(test.search)
		if test.expectedError != nil && err != nil {
			if err.Error() != test.expectedError.Error() {
				t.Fatalf("Expected error: %s Got error: %s", test.expectedError, err)
			}
			continue
		}
		if number != test.result {
			t.Fatalf("Number is not matching correctly: %v != %v", number, test.result)
		}
	}
}

func TestNumberService_DeleteNumber(t *testing.T) {
	tests := []struct {
		input         model.Numbers
		search        int
		expectedError error
	}{
		{
			model.Numbers{1},
			1,
			nil,
		},
		{
			model.Numbers{2},
			2,
			nil,
		},
		{
			model.Numbers{3},
			3,
			nil,
		},
		{
			model.Numbers{4},
			5,
			fmt.Errorf("Value not found: %v", 5),
		},
	}
	numRepository := repositories.NewNumberMemoryRepository()
	numService := NewNumberSerivce(numRepository)
	for _, test := range tests {
		numService.AddNumber(test.input)
		err := numService.DeleteNumber(test.search)
		if test.expectedError != nil && err != nil {
			if err.Error() != test.expectedError.Error() {
				t.Fatalf("Expected error: %s Got error: %s", test.expectedError, err)
			}
			continue
		}
	}
}
