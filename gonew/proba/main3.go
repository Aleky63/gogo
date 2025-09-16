package main3

import "fmt"

type Wallett struct {
	Cash int
}

// Pay implements Payer.
func (w *Wallett) Pay(int) error {
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

type Payerr interface {
	Pay(int) error
}

func Buyy(p Payerr) {
	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

func main3() {
	myWallet := &Wallett{Cash: 100}
	Buyy(myWallet)

	var myMoney Payerr
	myMoney = &Card{Balance: 100, Cardholder: "rvas"}
	Buyy(myMoney)

	myMoney = &ApplePay{Money: 9}

	Buyy(myMoney)

}
