package methods

import (
	"fmt"
	"math/rand"
)

type Paypal struct{}

func NewPaypal() Paypal {
	return Paypal{}
}

func (p Paypal) Pay(usd int) int {
	fmt.Println("Оплата через paypal!")
	fmt.Println("Размер оплаты:", usd, "usd")

	return rand.Int()
}
func (p Paypal) Cancel(id int) {
	fmt.Println("Отмена  paypal-операции! ID:", id)

}
