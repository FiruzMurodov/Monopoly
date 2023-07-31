package main

import (
	"fmt"
	"monopoly1/pkg"
	"strconv"
)

func main() {

	var count int
	fmt.Println("Сколько игроков начинают игру?")
	fmt.Scan(&count)
	players := []*pkg.Player{
		{},
	}
	for i := 1; i <= count; i++ {

		playersGaming := []*pkg.Player{
			{
				Name:     "Player" + strconv.Itoa(i),
				Money:    pkg.BeginingSum,
				Position: 0,
			},
		}
		players = append(players, playersGaming...)
	}

	for {
		for _, player := range players {
			if player.Money <= 0 {
				continue
			}
			rollDice := pkg.RandomDice()
			fmt.Printf("%s rolled a %d\n ", player.Name, rollDice)
			player.Move(rollDice)
			property := properties[player.Position]
			if property.Owner == nil {
				if player.Money >= property.Price {
					player.BuyProperty(property)
				} else if property.Owner != player {
					player.PayRent(property)
				}
				if player.Money <= 0 {
					fmt.Printf("%s is bankrupt.", player.Name)
					return
				}
			}

		}
	}

	//
	//var players = make([]int, count)
	//
	//for i := 0; i < len(players); i++ {
	//	pkg.Players{Position: i, Money: pkg.BeginingSum}
	//}
	//fmt.Println(players)
	//
	//var polyaMonopoly = make([]bool, 40)
	//var checkPlayer []bool
	//
	//for i := 0; players[i] >= 0; {
	//	khodPlayers := pkg.RandomDice(1, 6) +
	//		pkg.RandomDice(1, 6)
	//
	//	if polyaMonopoly[khodPlayers+1] == false {
	//
	//	} else {
	//
	//	}
	//}
}

var properties = []pkg.Property{
	{
		Name:  "First field",
		Price: 10,
		Rent:  5,
	},
	{
		Name:  "Second field",
		Price: 15,
		Rent:  10,
	},
	{
		Name:  "Third field",
		Price: 20,
		Rent:  15,
	},
	{
		Name:  "Fourth field",
		Price: 25,
		Rent:  20,
	},
	{
		Name:  "Fifth field",
		Price: 30,
		Rent:  25,
	},
	{
		Name:  "Sixth field",
		Price: 35,
		Rent:  30,
	},
	{
		Name:  "Seventh field",
		Price: 40,
		Rent:  35,
	},
	{
		Name:  "Eighth field",
		Price: 45,
		Rent:  40,
	},
	{
		Name:  "Ninth field",
		Price: 50,
		Rent:  45,
	},
	{
		Name:  "Tenth field",
		Price: 55,
		Rent:  50,
	},
	{
		Name:  "Eleventh field",
		Price: 60,
		Rent:  55,
	},
	{
		Name:  "Twelfth field",
		Price: 65,
		Rent:  55,
	},
	{
		Name:  "Thirteenth field",
		Price: 70,
		Rent:  65,
	},
	{
		Name:  "Fourteenth field",
		Price: 75,
		Rent:  70,
	},
	{
		Name:  "Fifteenth field",
		Price: 80,
		Rent:  75,
	},
}
