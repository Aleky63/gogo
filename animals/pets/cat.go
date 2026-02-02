package pets

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Cat struct {
	Name string
}

func (c *Cat) Eat(amount int) (int, error) {
	if amount > 4 {
		return 0, newError("Cat can't eat that much", nil)
	}
	return amount, nil
}
func (c *Cat) Walk() string {
	return "Cat is walking!!🐈‍⬛🐈‍⬛🐈‍⬛"
}
func (c *Cat) GetName() string {

	caser := cases.Title(language.English)

	return caser.String( c.Name)
}
