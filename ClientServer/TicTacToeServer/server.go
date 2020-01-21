package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

const address = ":8000"

var rooms = make([]*GameRoom, 0)
var mutex sync.Mutex

// connection handler (close)
func closeConnHandler(conn net.Conn) error {
	err := conn.Close()
	if err != nil {
		err = fmt.Errorf("error closing the connection")
		return err
	}
	return nil
}

func connectionHandler(conn net.Conn) {
	var readBuf = make([]byte, 100)
	conn.Read(readBuf)
	var gRoom *GameRoom = new(GameRoom)
	var player *Player = new(Player)
	mutex.Lock()
	// search of available room with ready host
	for _, curRoom := range rooms {
		if curRoom.isAvailable {
			gRoom = curRoom
			gRoom.playersCount++
			gRoom.isAvailable = false
			gRoom.pConn[conn] = true
			player.initPlayer(string(readBuf), "o", false, false, conn)
		}
	}
	mutex.Unlock()

	/*create new game room
	If the player variable was not initialized earlier (as a guest),
	it means that there were no available rooms.
	*/
	if player.name == "" {
		gRoom.initRoom("room" + strconv.Itoa(len(rooms)))
		gRoom.pConn[conn] = true
		mutex.Lock()
		rooms = append(rooms, gRoom)
		mutex.Unlock()
		// waiting opponent
		for gRoom.isAvailable == true {
		}
		player.initPlayer(string(readBuf), "x", true, true, conn)
	}
	gameSession(gRoom, player)

	if gRoom.pConn[player.conn] == true {
		gRoom.disconnChan <- struct{}{}
	}

	err := closeConnHandler(conn)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	fmt.Println("server is running")
	ln, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("error during listening", err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("error while accepting a connection", err)
		}
		go connectionHandler(conn)
	}
}
