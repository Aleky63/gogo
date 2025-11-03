package methods

import (
	"fmt"
	"math/rand"
)

type Bonus struct{}

func NewBonus() Bonus {
	return Bonus{}
}

func (b Bonus) Pay(usd int) int {
	fmt.Println("Оплата бонусами!")
	fmt.Println("Размер оплаты:", usd, "бонусов")

	return rand.Int()
}
func (b Bonus) Cancel(id int) {

	fmt.Println("Отмена  операции с бонусами! ID:", id)

}
