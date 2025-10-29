package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentsInfo  map[int]PaymentInfo
	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentMethod: paymentMethod,
	}
}

func (p PaymentModule) Pay(description string, usd int) int {
	return p.paymentMethod.Pay(usd)

}

func (p PaymentModule) Cancel(id int) {}

func (p PaymentModule) Info() {}

func (p PaymentModule) AllInfo() {}
