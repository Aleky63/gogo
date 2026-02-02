package pets

import (

	"fmt"
	// "golang.org/x/text/cases"
	// "golang.org/x/text/language"
)

type Eater interface {
	Eat(amount int) (int, error)
}
type Walker interface {
	Walk() string
}
type Named interface {
	GetName() string
}
type EaterWalker interface{
	Eater
	Walker
}






func newError(msg string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", msg, err)
	}
	return fmt.Errorf(msg)
}