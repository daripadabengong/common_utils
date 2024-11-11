package paginator

import (
	"fmt"
)

type Operator struct {
	value string
}

const (
	EqualsOperator      = "EQUALS"
	ContainsOperator    = "CONTAINS"
	GreaterThanOperator = "GREATER_THAN"
	LessThanOperator    = "LESS_THAN"
)

var validOperators = map[string]struct{}{
	EqualsOperator:      {},
	ContainsOperator:    {},
	GreaterThanOperator: {},
	LessThanOperator:    {},
}

// NewOperator is a factory function to create an Operator with validation
func NewOperator(value string) (Operator, error) {
	if _, exists := validOperators[value]; !exists {
		return Operator{}, fmt.Errorf("invalid operator: %s", value)
	}
	return Operator{value: value}, nil
}

// Value returns the string representation of the operator
func (o Operator) Value() string {
	return o.value
}

// Equals checks if two operators are the same
func (o Operator) Equals(other Operator) bool {
	return o.value == other.value
}
