package main

import (
	"4iar/board"
	"4iar/game"
	"4iar/player"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	numberOfRounds := flag.Int("n", 1, "numbers of rounds to play")
	flag.Parse()
	if *numberOfRounds < 1 {
		log.Printf("unable to play %d rounds\n", *numberOfRounds)
		os.Exit(1)
	}
	playerOne := player.NewRandomPlayer(board.PlayerOne)
	playerTwo := player.NewRandomPlayer(board.PlayerTwo)
	duel := game.NewGame(playerOne, playerTwo)
	outcome, err := duel.Play()
	if err != nil {
		log.Printf("play duel: %v\n", err)
		os.Exit(1)
	}
	if outcome == board.PlayerOneWins {
		fmt.Println("Player One Wins")
	} else if outcome == board.PlayerTwoWins {
		fmt.Println("Player Two Wins")
	} else if outcome == board.Tie {
		fmt.Println("Tied")
	} else {
		fmt.Println("Undecided")
	}
}
