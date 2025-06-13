package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	player := Player{}

	fmt.Println("Welcome to Terminal Roulette!!")
	time.Sleep(1 * time.Second)
	clearConsole()
	player.StartingMoney()

	for {
		player.PlaceBet()
		player.ColorOrNum()
		winningSlot := spinAndReveal()

		if player.slotBet.isColorBet == true && (winningSlot.Color == player.slotBet.colorChosen) {
			fmt.Printf("Congratulations you won $%d!\n", player.currentBet)
			player.WonBet()
			fmt.Printf("Current money: $%d\n", player.totalMoney)
		} else if player.slotBet.isNumberBet == true && (winningSlot.Number == player.slotBet.numberChosen) {
			fmt.Printf("Congratulations you won $%d!\n", player.currentBet)
			player.WonBet()
			fmt.Printf("Current money: $%d\n", player.totalMoney)
		} else {
			fmt.Printf("Oh no, you lost $%d!\n", player.currentBet)
			player.LostBet()
			fmt.Printf("Current money: $%d\n", player.totalMoney)
		}

		answer := PlayAgainPrompt()
		if answer == 1 {
			continue
		}
		break
	}

}

func ReadIntegerInput(prompt string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}

func PlayAgainPrompt() int {
	for {
		fmt.Println("\n\nPlay again?")
		fmt.Println("1. Yes")
		fmt.Println("2. No")

		answer, err := ReadIntegerInput("Enter Answer: ")
		if err != nil {
			fmt.Println("Invalid entry. Please try again.")
			continue
		} else if answer < 0 || answer > 2 {
			fmt.Println("Invalid entry. Please try again.")
			continue
		}
		return answer
	}
}
