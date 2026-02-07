package device

import "fmt"

type Pc struct {
	Cpu
}

func (p Pc) On() {
	fmt.Println("On compiter")
}
func (p Pc) Off() {
	fmt.Println("Off compiter")
}