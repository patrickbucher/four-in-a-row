package main

import (
	"4iar/player"
	"4iar/tournament"
	"flag"
	"fmt"
	"log"
)

func main() {
	numberOfRounds := flag.Int("n", 1, "number of rounds to play (with match and rematch)")
	flag.Parse()
	if *numberOfRounds < 1 {
		log.Fatalf("unable to play tournament with %d rounds", *numberOfRounds)
	}
	t := tournament.NewTournament()
	t.AddPlayer("Randy Random I.", player.NewRandomPlayer)
	t.AddPlayer("Randy Random II.", player.NewRandomPlayer)
	result, err := t.Play(*numberOfRounds)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
