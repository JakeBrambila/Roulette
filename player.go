package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bet struct {
	isNumberBet  bool
	isColorBet   bool
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
		fmt.Printf("How much do you wanna bet? ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid entry.")
			continue
		}
		input = strings.TrimSpace(input)
		bet, err := strconv.Atoi(input)
		if err != nil || bet < 0 || bet > p.totalMoney {
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
			p.slotBet.isColorBet = true
			p.slotBet.colorChosen = "R"
		} else if bet == 2 {
			p.slotBet.isColorBet = true
			p.slotBet.colorChosen = "B"
		} else if bet == 3 {
			p.slotBet.isNumberBet = true
			p.chooseNum()
		}
		break
	}
}

func (p *Player) chooseNum() {
	var numChosen int
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Choose a number between 1-36: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid entry please try again.")
			continue
		}
		input = strings.TrimSpace(input)
		num, error := strconv.Atoi(input)
		if error != nil || (num < 1 || num > 36) {
			fmt.Println("Invalid entry please try again.")
			continue
		}
		numChosen = num
		break
	}
	p.slotBet.numberChosen = numChosen
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

func (p *Player) ShowPlayerStats() {
	//shows how much you lost or won after the whole game is done
}
