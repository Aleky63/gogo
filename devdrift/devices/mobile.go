package device

import "fmt"

type Mobile struct {
	Cpu
}

func (m Mobile) On() {
	fmt.Println("On telephon")
}
func (m Mobile) Off() {
	fmt.Println("Off telephon")
}