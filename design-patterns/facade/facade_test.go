package facade

import "testing"

func TestFacadeAPI(t *testing.T) {
	api := NewApi()

	ret := api.Test()

	println(ret)
}
