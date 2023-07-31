package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

const BeginingSum = 1500

type Player struct {
	Name     string
	Money    int
	Position int
}

type Property struct {
	Name  string
	Price int
	Owner *Player
	Rent  int
}

func RandomDice() int {
	rand.Seed(time.Now().UnixNano())
	firstNumber := rand.Intn(6) + 1
	secondNumber := rand.Intn(6) + 1
	return firstNumber + secondNumber
}

func (p *Player) Move(steps int) {
	p.Position = (p.Position + steps) % 16
}

func (p *Player) BuyProperty(property Property) {
	p.Money -= property.Price
	property.Owner = p
	fmt.Printf("%s bought %s\n", p.Name, property.Name)
}
func (p *Player) PayRent(property Property) {
	p.Money -= property.Rent
	//enough money
	if property.Owner == nil {
		fmt.Println(property)
	} else {
		property.Owner.Money += property.Rent
	}
	fmt.Printf("%s paid rent to %s for %s\n", p.Name, property.Owner.Name, property.Name)
}
