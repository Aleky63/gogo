package pets



type Dog struct {
	Name string
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