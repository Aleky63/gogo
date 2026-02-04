package pets

type Dog struct {
	Animal
	Age      int
	Weight   int
	IsAsleep bool
}

func (d *Dog) Eat(amount int) (int, error) {
	
	if d.IsAsleep {
		return 0, &ActionError{Name: d.GetName(), Reason: "it is asleep"}
	}
	if amount > 10 {
		return 0, newError("Dog can't eat that much", nil)
	}
	return amount, nil
}
func (d *Dog) Walk() string {
	return "Dog is walking!!🐩 "
}