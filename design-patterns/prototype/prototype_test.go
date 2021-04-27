package prototype

import "testing"

var manager *Manager

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

func TestPrototype(t *testing.T) {

	if manager == nil {
		t.Fatal("manager not null")
	}

	t1 := manager.Get("t1")

	t2 := t1.Clone()

	if t1 == t2 {
		t.Fatal("error get clone not working")
	}
}

func init() {
	manager = NewPrototypeManager()

	manager.Set("t1", &Type1{
		name: "type1",
	})
}
