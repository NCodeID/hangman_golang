package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type User struct {
	nama  string
	score int
	tries int
}

func main() {
	var player User
	menu(&player)
	quiz := randomWord()
	player.tries = 6
	fmt.Println()
	gameStart(&player, quiz)
}

func menu(player *User) {
	enterName(player)
	clearTerminal()

}

func randomWord() string {
	keywords := readFile()
	quiz := rand.Intn(len(keywords))
	fmt.Println(quiz)
	return keywords[quiz]
}

func readFile() []string {
	file, err := os.Open("fruit.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []string{"null"}
	}
	defer file.Close()

	var fruits []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fruits = append(fruits, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return fruits
}

func enterName(player *User) {
	fmt.Print("Enter your username: ")
	fmt.Scan(&player.nama)
}

func gameStart(player *User, keyword string) {
	var state bool
	var answer string
	var temp = make([]string, 0)
	keyword = strings.Trim(strings.ReplaceAll(strings.ToLower(keyword), " ", ""), " ")

	for i := range keyword {
		temp = append(temp, "_")
		i++
	}
	for player.tries != 0 {
		state = false
		clearTerminal()
		fmt.Println("Welcome to hangman game")
		fmt.Println("User :", player.nama)
		fmt.Println("Score:", player.score)
		image(player.tries)
		for _, v := range temp {
			fmt.Print(v, " ")
		}
		fmt.Println("")
		validate := strings.Join(temp, "")
		if validate == keyword {
			fmt.Println("Congrats You're winning the game!")
			player.score++
			fmt.Println("Restart The Game (y/n): ")
			fmt.Scan(&answer)
			answer = strings.ToLower(answer)
			if answer == "yes" || answer == "y" {
				player.tries = 6
				gameStart(player, randomWord())
			} else {
				return
			}
			return

		}
		fmt.Print("Masukkan jawaban : ")
		fmt.Scanln(&answer)
		fmt.Println()
		answer = strings.ToLower(answer)
		for i, v := range keyword {
			if answer == string(v) {
				temp[i] = answer
				state = true
			}
		}
		if !state {
			player.tries--
		}

	}
	clearTerminal()
	image(player.tries)
	fmt.Println("You lost the game")
}
func clearTerminal() {
	if runtime.GOOS == "windows" {
		c := exec.Command("cls")
		c.Stdout = os.Stdout
		c.Run()
	} else {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	}
}
func image(state int) {
	var gambar string
	switch state {
	case 6:
		gambar =
			`
  +---+
      |
      |
      |
      |
      |
=========
		
		`
	case 5:
		gambar =
			`
  +---+
  |   |
  O   |
      |
      |
      |
=========
		
		`
	case 4:
		gambar =
			`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========
		
		`
	case 3:
		gambar =
			`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========
		
		`
	case 2:
		gambar =
			`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========
		
		`
	case 1:
		gambar =
			`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========
		
		`
	case 0:
		gambar =
			`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========
	`
	default:
		gambar =
			`
  +---+
      |
      |
      |
      |
      |
=========
		
		`
	}

	fmt.Println(gambar)
}
