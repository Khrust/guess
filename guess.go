package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var multiple int = 100
var attemptsCount int = 10
var secretNumber int

func initGame() {
	rand.Seed(time.Now().UnixNano())
	secretNumber = rand.Intn(multiple) + 1
}

func main() {

	initGame()

	fmt.Printf("Guess the number... ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	fmt.Println(input)

	fmt.Println(secretNumber)
}
