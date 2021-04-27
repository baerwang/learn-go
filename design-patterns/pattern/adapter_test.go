package pattern

import "testing"

func TestAdapter(t *testing.T) {
	adpatee := NewAdaptee()

	target := NewAdapter(adpatee)

	res := target.Request()
	if res != "adpatee method" {
		t.Fatalf("expect :%s,actual:%s", "adaptee method", res)
	}

}
