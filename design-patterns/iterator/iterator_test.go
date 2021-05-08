package iterator

import "testing"

func TestIterator(t *testing.T) {

	IteratorsPrint(NewNumbers(1, 10).Iterator())

}
