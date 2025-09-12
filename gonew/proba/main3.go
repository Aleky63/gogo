package main

import "fmt"

type Wallet struct {
	Cash int
}

// Pay implements Payer.
func (w *Wallet) Pay(int) error {
	panic("unimplemented")
}

type Card struct {
	Balance    int
	ValidUntil string
	Cardholder string
	CVV        string
	Number     string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Мало денег card")
	}
	c.Balance -= amount
	return nil
}

type ApplePay struct {
	Money   int
	AppleID string
}

func (a *ApplePay) Pay(amount int) error {

	if a.Money < amount {
		return fmt.Errorf("Мало денег akk")
	}
	a.Money -= amount
	return nil
}

type Payer interface {
	Pay(int) error
}

func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

func main3() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	var myMoney Payer
	myMoney = &Card{Balance: 100, Cardholder: "rvas"}
	Buy(myMoney)

	myMoney = &ApplePay{Money: 9}

	Buy(myMoney)

}
