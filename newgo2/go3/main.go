package main

import (
	"go2/payments"
	"go2/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {
	method := methods.NewCrypto()
	paymentModule := payments.NewPaymentModule(method)
	paymentModule.Pay("Бургер", 5)
	paymentModule.Pay("Телефон", 598)
	paymentModule.Pay("Игрушка", 49)

	allInfo := paymentModule.AllInfo()
	pp.Println("Все оплаты наши:", allInfo)
}
