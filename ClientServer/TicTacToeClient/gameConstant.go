package main

//GameState describes the player's in-game state
type GameState int

const (
	//Move - player's move
	Move GameState = iota
	//Wait - waiting for the opponent's move
	Wait
	//GameOver - game over
	GameOver
)

//GameResult describes various options for ending the game
type GameResult int

const (
	//Win ...
	Win GameResult = iota
	//Lose ...
	Lose
	//Draw ...
	Draw
)