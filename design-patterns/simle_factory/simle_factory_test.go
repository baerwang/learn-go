package simle_factory

import "testing"

func TestType1(t *testing.T) {
	api := NewApi(1)

	s := api.Say("tom")

	if s != "hi,tom" {
		t.Fatal("type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api := NewApi(2)

	s := api.Say("wxx")

	if s != "hello,wxx" {
		t.Fatal("typ2 test fail")
	}

}
