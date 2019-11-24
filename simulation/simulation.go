package main

import (
	"4iar/board"
	"4iar/game"
	"4iar/player"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
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
	playerOneWins, playerTwoWins, ties, undecided := 0, 0, 0, 0
	output := *numberOfRounds == 1
	ch := make(chan board.Outcome)
	var wg sync.WaitGroup
	for i := 0; i < *numberOfRounds; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			duel := game.NewGame(playerOne, playerTwo)
			outcome, err := duel.Play(output)
			if err != nil {
				log.Printf("play duel: %v\n", err)
				return
			}
			ch <- outcome
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for outcome := range ch {
		if outcome == board.PlayerOneWins {
			if output {
				fmt.Println("Player One Wins")
			}
			playerOneWins++
		} else if outcome == board.PlayerTwoWins {
			if output {
				fmt.Println("Player Two Wins")
			}
			playerTwoWins++
		} else if outcome == board.Tie {
			if output {
				fmt.Println("Tied")
			}
			ties++
		} else {
			if output {
				fmt.Println("Undecided")
			}
			undecided++
		}
	}
	if !output {
		fmt.Printf("Player One Wins: %8d\n", playerOneWins)
		fmt.Printf("Player Two Wins: %8d\n", playerTwoWins)
		fmt.Printf("Ties:            %8d\n", ties)
		fmt.Printf("Undecided:       %8d\n", undecided)
	}
}
