package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
)

// server address
const addr = ""

// show the current version of the "game field"
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

// decode the byte representation of the "game field"
func decode(encodeBuf []byte) ([9]string, error) {
	var tempArr [9]string
	var byteBuffer bytes.Buffer
	byteBuffer.Write(encodeBuf)
	dec := gob.NewDecoder(&byteBuffer)
	if err := dec.Decode(&tempArr); err != nil {
		decodeEr := fmt.Errorf("it is not possible to decode the buffer")
		return tempArr, decodeEr
	}

	//if a "zero" field is received, then the opponent has been disconnected
	if tempArr[0] == "0" {
		err := fmt.Errorf("oops, your opponent is disconnected. You win")
		return tempArr, err
	}

	return tempArr, nil
}

func main() {
	var name string
	var fld = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var numCell int
	fmt.Println("Welcom to \"TicTacToe\"")
	fmt.Println("Please, enter your name")
	fmt.Scan(&name)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("unable to connect to the server, try again later")
		return
	}
	// set player name
	conn.Write([]byte(name))
	// main loop
	for {
		buf := make([]byte, 1)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("could not get information from the server, try connecting later")
			return
		} else if state, _ := strconv.Atoi(string(buf)); GameState(state) == Move {
			fmt.Println(name, "is move now")
			fmt.Println("Please, enter the num of cell, 1 <= num <= 9")
			fmt.Scan(&numCell)
			strRepr := strconv.Itoa(numCell)
			conn.Write([]byte(strRepr))
			fmt.Println("You made a move, now wait for your opponent")
		} else if state, _ := strconv.Atoi(string(buf)); GameState(state) == Wait {
			var decodeErr error
			fmt.Println("Your opponent is now making a move, wait please")
			fieldBuf := make([]byte, 38)
			conn.Read(fieldBuf)
			fld, decodeErr = decode(fieldBuf)
			if decodeErr != nil {
				fmt.Println(decodeErr)
				return
			}
			showGameField(fld)
		} else if state, _ := strconv.Atoi(string(buf)); GameState(state) == GameOver {
			resultBuf := make([]byte, 1)
			conn.Read(resultBuf)
			switch res, _ := strconv.Atoi(string(resultBuf)); GameResult(res) {
			case Win:
				fmt.Println("Congratulations, You have won")
			case Lose:
				fmt.Println("Unfortunately, you lost")
			case Draw:
				fmt.Println("Game is draw")
			}
			conn.Close()
			break
		}
	}
}
