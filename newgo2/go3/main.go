package main

import (
	"go3/payments"
	"go3/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {
	method := methods.NewCrypto()
	paymentModule := payments.NewPaymentModule(method)
	paymentModule.Pay("Бургер", 5)

	idPhone := paymentModule.Pay("Телефон", 598)

	idGame := paymentModule.Pay("Игрушка", 49)

	paymentModule.Cancel((idPhone))

	allInfo := paymentModule.AllInfo()
	pp.Println("Все оплаты наши:", allInfo)
	gameInfo := paymentModule.Info((idGame))
	pp.Println("Game info:", gameInfo)
}
