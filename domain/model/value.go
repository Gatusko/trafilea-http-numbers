package model

type Value struct {
	Number int `json:"number"`
	Val    any `json:"value"` // Here we can return an Int or String based in our information
}

func NewValue(num int) Value {
	val := Value{
		Number: num,
	}
	switch {
	case num%3 == 0 && num%5 == 0:
		val.Val = "Type 3"
	case num%5 == 0:
		val.Val = "Type 2"
	case num%3 == 0:
		val.Val = "Type 1"
	default:
		val.Val = num
	}
	return val
}
