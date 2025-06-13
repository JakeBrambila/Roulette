package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type Slot struct {
	Number int
	Color  string // "R", "B", or "G"
}

const (
	Red   = "\033[31m"
	Black = "\033[30m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

var wheel = []Slot{
	{0, "G"},
	{32, "R"}, {15, "B"}, {19, "R"}, {4, "B"}, {21, "R"},
	{2, "B"}, {25, "R"}, {17, "B"}, {34, "R"}, {6, "B"},
	{27, "R"}, {13, "B"}, {36, "R"}, {11, "B"}, {30, "R"},
	{8, "B"}, {23, "R"}, {10, "B"}, {5, "R"}, {24, "B"},
	{16, "R"}, {33, "B"}, {1, "R"}, {20, "B"}, {14, "R"},
	{31, "B"}, {9, "R"}, {22, "B"}, {18, "R"}, {29, "B"},
	{7, "R"}, {28, "B"}, {12, "R"}, {35, "B"}, {3, "R"},
	{26, "B"},
}

func spinAndReveal() Slot {
	spinLength := 30
	stripSize := 11 // Display window

	// Start at a random point in the wheel
	start := rand.Intn(len(wheel))

	for i := 0; i < spinLength; i++ {
		clearConsole()
		offset := (start + i) % len(wheel)
		showStrip(offset, stripSize)
		time.Sleep(100 * time.Millisecond)
	}

	// Final result (center of strip)
	clearConsole()
	finalOffset := (start + spinLength) % len(wheel)
	showStrip(finalOffset, stripSize)
	centerIndex := stripSize / 2
	winner := wheel[(finalOffset+centerIndex)%len(wheel)]

	fmt.Printf("\n🎯 Winning Number: %s%d%s (%s)\n", colorCode(winner.Color), winner.Number, Reset, winner.Color)
	return Slot{winner.Number, winner.Color}
}

func showStrip(start int, size int) {
	fmt.Println("╔═══════════════════════════════════════════════════════════════════╗")
	fmt.Print("║ ")
	for i := 0; i < size; i++ {
		slot := wheel[(start+i)%len(wheel)]
		colored := fmt.Sprintf("%s%2d%s%s", colorCode(slot.Color), slot.Number, slot.Color, Reset)
		if i == size/2 {
			fmt.Printf("[%s] ", colored) // center highlight
		} else {
			fmt.Printf(" %s  ", colored)
		}
	}
	fmt.Println("║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════════╝")
}

func colorCode(c string) string {
	switch c {
	case "R":
		return Red
	case "B":
		return Black
	case "G":
		return Green
	default:
		return Reset
	}
}

func clearConsole() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") // Windows clear command
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin": // macOS
		cmd := exec.Command("clear") // macOS clear command
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux": // Linux
		fmt.Print("\033[H\033[2J") // ANSI escape sequence for clearing the terminal
	default: // Fallback for other Unix-based systems
		fmt.Print("\033[H\033[2J") // ANSI escape sequence for clearing the terminal
	}
}
