package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var hangmanDictionary = []string{
	"affix",
	"exodus",
	"disavow",
	"jockey",
	"kiosk",
	"pneumonia",
	"paralysis",
	"pshaw",
	"kayak",
	"cobweb",
	"avenue",
	"lymph",
}

var reader = bufio.NewReader(os.Stdin)
var inputSlice []string
var input string
var randSource = rand.NewSource(time.Now().UnixNano())
var randGen = rand.New(randSource)
var hangmanState int
var randomNo = randGen.Intn(len(hangmanDictionary))
var wordToBeGuessed = hangmanDictionary[randomNo]

func main() {
	var lettersGuessed = map[rune]bool{}
	lettersGuessed[rune(wordToBeGuessed[0])] = true
	lettersGuessed[rune(wordToBeGuessed[len(wordToBeGuessed)-1])] = true
	fmt.Println("You can enter 'hint' to use the hint ONCE!")
	var hintNum int
	for {
		printState(wordToBeGuessed, lettersGuessed, hangmanState)
		input, _ = getUserInput()
		input = strings.ToLower(input)
		if input == string(wordToBeGuessed[0]) || input == string(wordToBeGuessed[len(wordToBeGuessed)-1]) {
			fmt.Println("letter has been revealed already, try another unrevealed letter!")
			continue
		}
		if input == "hint" && hintNum < 1 {
			activateHint()
			hintNum++
			continue
		} else if input == "hint" && hintNum >= 1 {
			fmt.Println("Hint already used! ")
			continue
		} else if len(input) != 1 || !strings.ContainsAny(input, "abcdefghijklmnopqrstuvwxyz") {
			err := errors.New("invalid input, please input just a letter")
			fmt.Println(err)
			continue
		}
		if hasInputBeenEntered(inputSlice) {
			err := errors.New("input entered already, enter a new word")
			fmt.Println(err)
			continue
		}
		if strings.Contains(wordToBeGuessed, input) {
			lettersGuessed[rune(input[0])] = true
		} else {
			hangmanState++
		}
		if hangmanState >= 8 {
			printState(wordToBeGuessed, lettersGuessed, hangmanState)
			fmt.Println("Game Over..., You Lose")
			return
		}
		if isWordGuessed(wordToBeGuessed, lettersGuessed) {
			printState(wordToBeGuessed, lettersGuessed, hangmanState)
			fmt.Println("Game Over..., You Win!")
			return
		}
	}
}

func printState(wordToBeGuessed string, lettersGuessed map[rune]bool, hangmanState int) {
	for _, letter := range wordToBeGuessed {
		if lettersGuessed[letter] {
			fmt.Print(string(letter))
		} else {
			fmt.Print("_")
		}
		fmt.Print(" ")
	}
	fmt.Printf("\n" + "\n")
	fmt.Println(drawHangman(hangmanState))
}

func drawHangman(hangmanState int) string {
	bytesRead, err := os.ReadFile("states/figure" + fmt.Sprint(hangmanState))
	if err != nil {
		panic(err)
	}
	return string(bytesRead)
}

func getUserInput() (string, []string) {
	fmt.Print(">> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	inputSlice = append(inputSlice, input)
	return input, inputSlice
}

func isWordGuessed(wordToBeGuessed string, lettersGuessed map[rune]bool) bool {
	for _, v := range wordToBeGuessed {
		if lettersGuessed[v] {
			continue
		} else {
			return false
		}
	}
	return true
}

func hasInputBeenEntered(inputSlice []string) bool {
	var ret bool
	for i := 0; i < len(inputSlice)-1; i++ {
		if strings.EqualFold(input, inputSlice[i]) {
			ret = true
			break
		} else {
			ret = false
		}
	}
	return ret
}

func activateHint() {
	fmt.Print("Your Hint: ")
	rand.Seed(time.Now().UnixNano())
	randomNo := rand.Intn(len(wordToBeGuessed)-2) + 1
	fmt.Printf(string(wordToBeGuessed[randomNo]) + "\n")
}
