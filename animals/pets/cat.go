package pets

type Cat struct {
	Animal
	Age      int
	IsAsleep bool
}

func (c *Cat) Eat(amount int) (int, error) {
	
	if c.IsAsleep {
		return 0, &ActionError{Name: c.GetName(), Reason: "it is asleep"}
	}

	if amount > 5 {
		return 0, newError("Cat can't eat that much", nil)
	}
	return amount, nil
}
func (c *Cat) Walk() string {
	return "Cat is walking!!🐈‍⬛🐈‍⬛"
}
