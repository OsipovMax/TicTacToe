package main

import "net"

/*Player is a representation of one of the players.
*name - player name
*gameSymbol - symbol of the player ("x or o")
*isMove - status flag of the player ("move or wait opponent")
*isHost - host or guest of the game room
 */
type Player struct {
	name       string
	gameSymbol string
	isMove     bool
	isHost     bool
	conn       net.Conn
}

func (player *Player) initPlayer(pName string, symb string, move bool, host bool, c net.Conn) {
	player.name = pName
	player.gameSymbol = symb
	player.isMove = move
	player.isHost = host
	player.conn = c
}

/*GameRoom is a representation of a virtual room that unites a couple of players.
*roomTitle - title of the game room
*pConn - slice that stores the connection of players (host and guest)
*updateField - сhannel for transmitting data about the playing field
*
*isAvailable - flag access to the games room (if playerCount = 2 game room is not available)
*playerCount - number of players into game room
 */
type GameRoom struct {
	roomTitle    string
	pConn        map[net.Conn]bool
	updateField  chan [9]string
	disconnChan  chan struct{}
	isAvailable  bool
	playersCount int
}

func (room *GameRoom) initRoom(title string) {
	room.roomTitle = title
	room.pConn = make(map[net.Conn]bool, 2)
	room.updateField = make(chan [9]string)
	room.disconnChan = make(chan struct{})
	room.isAvailable = true
	room.playersCount = 1
}

//GameField ...
type GameField struct {
	generalSize int
	field       [9]string
}

func (gameField *GameField) initGameField() {
	gameField.generalSize = 9
	gameField.field = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
}
