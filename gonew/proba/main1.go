package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

func (p *Person) SetName(name string) {
	p.Name = name
}

type MySlice []int

func (sl *MySlice) Add(val int) {
	*sl = append(*sl, val)
}
func (sl *MySlice) Count() int {
	return len(*sl)
}

type Account struct {
	Id   int
	Name string
	Person
}

func main1() {

	pers := Person{1, "Vasiliy"}
	pers.SetName(("Vasiliy Romanov"))
	fmt.Println(pers)
	fmt.Printf("updated person: %#v\n", pers)

	var acc Account = Account{
		Id:   3,
		Name: "rrrrrrrrr",
		Person: Person{
			Id:   5,
			Name: "ASSSSSSS",
		},
	}
	acc.SetName("ROOOOOOO")
	fmt.Printf(" %#v\n", acc)
}
