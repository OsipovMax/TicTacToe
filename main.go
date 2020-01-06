package main

import (
	"fmt"
)

// Player information
type Player struct {
	playerName string
	gameSymbol string
}

func showGameField(field [9]string) {
	for ind, cell := range field {
		if ind == 3 || ind == 6 {
			fmt.Print("|")
			fmt.Printf("\n")
		}
		fmt.Printf("|%7s", cell)
	}
	fmt.Print("|\n")
}

func gameOver(field [9]string, pl1 Player, pl2 Player) bool {
	result := gameLogic(field)
	if result == "x" {
		fmt.Println("Player ", pl1.playerName, " won")
		return true
	}
	if result == "o" {
		fmt.Println("Player ", pl2.playerName, " won")
		return true
	}
	return false
}

func checkMap(moveCombination map[string]int) (bool, string) {
	for key, val := range moveCombination {
		if val == 3 {
			return true, key
		}
	}
	return false, ""
}
func gameLogic(field [9]string) string {
	m := make(map[string]int)
	//Horizontal check
	for i := 0; i < len(field); i += 3 {
		for j := 0; j < len(field)/3; j++ {
			m[field[i+j]]++
		}
		isSuccess, playerSign := checkMap(m)
		if isSuccess {
			return playerSign
		}
		m = make(map[string]int)
	}
	//Vertical check
	for j := 0; j < len(field)/3; j++ {
		for i := 0; i < len(field); i += 3 {
			m[field[j+i]]++
		}
		isSuccess, playerSign := checkMap(m)
		if isSuccess {
			return playerSign
		}
		m = make(map[string]int)
	}

	// Diagonal check
	for i := 0; i < len(field)/3; i++ {
		m[field[i*3+i]]++
	}
	isSuccess, playerSign := checkMap(m)
	if isSuccess {
		return playerSign
	}
	m = make(map[string]int)

	for i := 0; i < len(field)/3; i++ {
		m[field[(i+1)*3-(i+1)]]++
	}
	isSuccess, playerSign = checkMap(m)
	if isSuccess {
		return playerSign
	}

	return ""
}

func playersMove(field *[9]string, pl Player) {
	var numCell int
	for numCell < 1 || numCell > 9 {
		fmt.Println("Please slect number of cell")
		fmt.Scanf("%d\n", &numCell)
	}
	field[numCell-1] = pl.gameSymbol
	showGameField(*field)
}

func main() {
	var inputPlayerInfromation string
	gameRound := 1
	fmt.Println("Hello, it is \"TicTacToe\" game")
	fmt.Println("Player 1, please enter your name. Attention your sybol is \"x\"")
	fmt.Scanln(&inputPlayerInfromation)
	player1 := Player{playerName: inputPlayerInfromation, gameSymbol: "x"}
	fmt.Println("Player 2, please enter your name. Attention your sybol is \"o\"")
	fmt.Scanln(&inputPlayerInfromation)
	player2 := Player{playerName: inputPlayerInfromation, gameSymbol: "o"}
	field := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	fmt.Println("This is Your playing field")
	showGameField(field)
	fmt.Println("Let`s start the game, ", player1.playerName, " make the first move")
	for {
		if gameOver(field, player1, player2) {
			break
		}
		if gameRound%2 != 0 {
			fmt.Println(player1.playerName, "your move")
			playersMove(&field, player1)
		} else {
			fmt.Println(player2.playerName, "your move")
			playersMove(&field, player2)
		}
		gameRound++
	}
}
