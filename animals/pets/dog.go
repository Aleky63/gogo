package pets



type Dog struct {
	Animal
	Age int
	Weight int
	IsAsleep bool
}

func (c *Dog) Eat(amount int) (int, error) {
	if amount > 10 {
		return 0, newError("Dog can't eat that much", nil)
	}
	return amount, nil
}
func (d *Dog) Walk() string {
	return "Dog is walking!!🐩 🐩 🐩"
}