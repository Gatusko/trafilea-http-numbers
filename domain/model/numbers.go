package model

type Numbers struct {
	Value int `json:"number"`
}

func NewNumber() Numbers {
	return Numbers{}
}

// Function for validate in the future
func (n *Numbers) Validate() error {
	return nil
}
