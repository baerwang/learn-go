package observer

import (
	"strconv"
	"testing"
)

func TestObserver(t *testing.T) {
	subject := NewSubject()

	for i := 0; i < 3; i++ {
		subject.Attach(NewReader("reader" + strconv.Itoa(i)))
	}

	subject.UpdateContext("observer model")
}
