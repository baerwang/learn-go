package factory_method

import (
	"fmt"
	"testing"
)

func compute(factory OperatorFactory, a int, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)

	return op.Result()
}

func TestOperator(t *testing.T) {

	var factory OperatorFactory

	factory = PlusOperatorFactory{}

	fmt.Println(compute(factory, 1, 2))

	factory = MinusOperatorFactory{}

	fmt.Println(compute(factory, 100, 2))

}
