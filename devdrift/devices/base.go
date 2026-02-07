package device

type Device interface {
	On()
	Off()
	
}

type Cpu struct {
	Prod string
	Core int
}