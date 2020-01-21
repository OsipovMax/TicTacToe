package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
)

// write/read mode
type mode int

const (
	//write
	wr mode = iota
	//read
	rd
)

func encode(fieldArr [9]string) ([]byte, error) {
	var byteBuffer bytes.Buffer
	encodeBuf := make([]byte, 38)
	enc := gob.NewEncoder(&byteBuffer)
	if err := enc.Encode(fieldArr); err != nil {
		encodeErr := fmt.Errorf("encoding errror")
		return encodeBuf, encodeErr
	}
	byteBuffer.Read(encodeBuf)
	return encodeBuf, nil
}

func messaging(conn net.Conn, connMode mode, sndRcvBuffer []byte, errString string) error {
	if connMode == wr {
		_, writeErr := conn.Write(sndRcvBuffer)
		if writeErr != nil {
			writeErr = fmt.Errorf(errString)
			return writeErr
		}
	} else {
		_, readErr := conn.Read(sndRcvBuffer)
		if readErr != nil {
			readErr = fmt.Errorf(errString)
			return readErr
		}
	}
	return nil
}

func gameSession(room *GameRoom, player *Player) error {
	recvBuf := make([]byte, 1)
	var gameField = new(GameField)
	gameField.initGameField()
	for {
		if over, symb := gameOver(gameField.field); over == true {
			err := messaging(player.conn, wr, []byte(strconv.Itoa(int(GameOver))), "write error state - \"game over\"")
			if err != nil {
				return err
			}
			//first check for a draw
			if symb == "-" {
				err := messaging(player.conn, wr, []byte(strconv.Itoa(int(Draw))), "write error draw mark")
				if err != nil {
					return err
				}
				break
			}
			//win
			if symb == player.gameSymbol {
				err := messaging(player.conn, wr, []byte(strconv.Itoa(int(Win))), "write error win mark")
				if err != nil {
					return err
				}
			} else { //lose
				err := messaging(player.conn, wr, []byte(strconv.Itoa(int(Lose))), "write error lose mark")
				if err != nil {
					return err
				}
			}
			room.pConn[player.conn] = false
			break
		}
		if player.isMove == true {
			//send mark "move"
			err := messaging(player.conn, wr, []byte(strconv.Itoa(int(Move))), "write error state \"move\"")
			if err != nil {
				return err
			}
			//read cell num
			err = messaging(player.conn, rd, recvBuf, "read error cell num")
			if err != nil {
				return err
			}
			numCell, err := strconv.Atoi(string(recvBuf))
			if err != nil || numCell == 0 {
				err := fmt.Errorf("error receiving the cell num")
				return err
			}
			gameField.field[numCell-1] = player.gameSymbol
			player.isMove = false
			room.updateField <- gameField.field
		} else if player.isMove == false {
			err := messaging(player.conn, wr, []byte(strconv.Itoa(int(Wait))), "write error state \"wait\"")
			if err != nil {
				return err
			}
			select {
			case gameField.field = <-room.updateField:
				{
					encodeField, err := encode(gameField.field)
					if err != nil {
						return err
					}
					err = messaging(player.conn, wr, encodeField, "write error updated game field")
					if err != nil {
						return err
					}
					player.isMove = true
				}
			case <-room.disconnChan:
				{
					//zero := [...]string{"0", "0", "0", "0", "0", "0", "0", "0", "0"}
					zeroField, err := encode([...]string{"0", "0", "0", "0", "0", "0", "0", "0", "0"})
					if err != nil {
						return err
					}

					err = messaging(player.conn, wr, zeroField, "write error \"zero\" field")
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}
