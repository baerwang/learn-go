package command

import "testing"

func TestCommand(t *testing.T) {
	mb := &MotherBoard{}

	startCommand := NewStartCommand(mb)
	rebootCommand := NewRebootCommand(mb)

	box1 := NewBox(startCommand, rebootCommand)
	box1.PressButton1()
	box1.PressButton2()
}
