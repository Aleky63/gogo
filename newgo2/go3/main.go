package main

import (
	"go3/payments"
	"go3/payments/methods"

	"github.com/k0kubun/pp"
)

// ------------------------

func main() {
	method := methods.NewBank()
	paymentModule := payments.NewPaymentModule(method)
	paymentModule.Pay("бургер", 5)

	idPhone := paymentModule.Pay("телефон", 666)
	idGame := paymentModule.Pay("игрушка", 120)

	paymentModule.Cancel(idPhone)

	allInfo := paymentModule.AllInfo()
	pp.Println("Все наши оплаты", allInfo)

	gameInfo := paymentModule.Info(idGame)
	pp.Println("Game info", gameInfo)
}
