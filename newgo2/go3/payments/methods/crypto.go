package methods

import (
	"fmt"
	"math/rand"
)

type Crypto struct{}

func NewCrypto() Crypto {
	return Crypto{}
}

func (c Crypto) Pay(usd int) int {
	fmt.Println("Оплата криптой!")
	fmt.Println("Размер оплаты криптой:", usd, "USDT")

	return rand.Int()
}
func (c Crypto) Cancel(id int) {
	fmt.Println("Отмена криптооперации! ID:", id)

}
