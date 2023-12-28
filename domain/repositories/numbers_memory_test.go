package repositories

import (
	"fmt"
	"github.com/Gatusko/trafilea-http-numbers/domain/model"
	"testing"
)

func TestNumbersMemory_Add(t *testing.T) {
	tests := []struct {
		input         model.Numbers
		value         model.Value
		result        model.Value
		expectedError error
	}{
		{
			model.Numbers{1},
			model.Value{1},
			model.Value{1},
			nil,
		},
		{
			model.Numbers{5},
			model.Value{"test"},
			model.Value{"test"},
			nil,
		},
	}
	numRepository := NewNumberMemoryRepository()
	for _, test := range tests {
		number, err := numRepository.Add(test.input, test.value)
		if err != test.expectedError {
			t.Fatalf("Expected error: %s", test.expectedError)
		}
		if number != test.result {
			t.Fatalf("Number is not matching correctly: %v != %v", number, test.result)
		}
	}
}

// This cover Get/Delete
func TestNumbersMemory_DeleteGet(t *testing.T) {
	tests := []struct {
		input         model.Numbers
		value         model.Value
		deleteNumber  int
		expectedError error
	}{
		{
			model.Numbers{1},
			model.Value{1},
			5,
			fmt.Errorf("Value not found: %v", 5),
		},
		{
			model.Numbers{3},
			model.Value{"test"},
			3,
			nil,
		},
	}
	numRepository := NewNumberMemoryRepository()
	for _, test := range tests {
		numRepository.Add(test.input, test.value)
		err := numRepository.Delete(test.deleteNumber)
		if test.expectedError != nil && err != nil {
			if err.Error() != test.expectedError.Error() {
				t.Fatalf("Expected error: %s", test.expectedError)
			}
		}
	}
}
