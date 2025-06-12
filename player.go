package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bet struct {
	numberBet    bool
	colorBet     bool
	colorChosen  string
	numberChosen int
}

type Player struct {
	startingMoney int
	totalMoney    int
	currentBet    int
	slotBet       Bet
}

func (p *Player) WonBet() {
	p.totalMoney += p.currentBet
	p.currentBet = 0
}

func (p *Player) LostBet() {
	p.totalMoney -= p.currentBet
	p.currentBet = 0
}

func (p *Player) Tie() {
	p.currentBet = 0
}

func (p *Player) StartingMoney() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("How much money would you like to start with? ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid entry.")
			continue
		}
		input = strings.TrimSpace(input)
		money, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid entry.")
			continue
		}
		p.totalMoney = money
		p.startingMoney = money
		break
	}
}

func (p *Player) PlaceBet() {
	for {
		fmt.Println("How much do you wanna bet? ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid entry.")
			continue
		}
		input = strings.TrimSpace(input)
		bet, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid entry.")
			continue
		}
		p.currentBet = bet
		break
	}
}

func (p *Player) ColorOrNum() {
	fmt.Print(colorBetMenu)
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid entry.")
			continue
		}
		input = strings.TrimSpace(input)
		bet, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid entry.")
			continue
		}
		if bet == 1 {
			//bet is red
		} else if bet == 2 {
			//bet is on black
		} else if bet == 3 {
			//bet is a number
			//needs more logic here
		}

		break
	}
}

func (p *Player) CalculateEarnings() {
	if p.startingMoney > p.totalMoney {
		loss := p.startingMoney - p.totalMoney
		fmt.Printf("You lost $%d\n", loss)
		fmt.Println("Better luck next time!")
		return
	} else if p.startingMoney == p.totalMoney {
		fmt.Println("You lost no money! Congratulations!")
		fmt.Println("Maybe win some next time loser.")
		return
	}
	won := p.totalMoney - p.startingMoney
	fmt.Printf("You won $%d\n", won)
	if won < p.startingMoney {
		fmt.Println("You didn't win that much so keep playing.")
	} else if won < (p.startingMoney * 2) {
		fmt.Println("You won a decent amount good job.")
	} else {
		fmt.Println("You are a true gambler. I bow down to you master.")
	}
}
