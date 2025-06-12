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

	spinAndReveal()
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
