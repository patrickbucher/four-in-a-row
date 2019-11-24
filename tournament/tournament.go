// Package tournament contains the logic to let a number of different players
// play against one another multiple times.
package tournament

import (
	"4iar/board"
	"4iar/game"
	"4iar/player"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
)

// PlayerSpawnFunc is a function that creates a new player with the given
// field.
type PlayerSpawnFunc func(board.Field) *player.Player

// Tournament is a set of named players, which are created using their
// PlayerSpawnFunc.
type Tournament map[string]PlayerSpawnFunc

// NewTournament creates a new, empty tournament, i.e. without players.
func NewTournament() *Tournament {
	t := Tournament(make(map[string]PlayerSpawnFunc, 0))
	return &t
}

// AddPlayer adds a new player to the tournament. If a player with name was
// already added to the tournament before, an error is returned.
func (t *Tournament) AddPlayer(name string, spawnFunc PlayerSpawnFunc) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("empty name for player is not allowed")
	}
	if spawnFunc == nil {
		return errors.New("spawnFunc must not be nil")
	}
	if _, ok := (*t)[name]; ok {
		return fmt.Errorf("a player with name='%s' was added before", name)
	}
	(*t)[name] = spawnFunc
	return nil
}

// Player is a player in the context of a tournament (potential player, with
// name and spawn func).
type Player struct {
	Name      string
	SpawnFunc PlayerSpawnFunc
}

// Pairing is match pairing of two created players and their names.
type Pairing struct {
	PlayerOne     *player.Player
	PlayerOneName string
	PlayerTwo     *player.Player
	PlayerTwoName string
}

// Play plays the given number of rounds and returns the resulting tournament
// statistics. Every player is paired up twice with each other player of the
// tournament in flipped order to compensate for a possible first-mover
// advantage. If less than two players have been added to the tournament, an
// error is returned.
func (t *Tournament) Play(rounds int) (Result, error) {
	if len(*t) < 2 {
		return nil, errors.New("unable to play a tournament with less than two players")
	}
	pairings := pairUp(t)
	stats := make(map[string]*PlayerStatistics, 0)
	for name := range *t {
		ps := PlayerStatistics{name, 0, 0, 0, 0, 0}
		stats[name] = &ps
	}
	var wg sync.WaitGroup
	deltaStatChan := make(chan PlayerStatistics)
	for r := 0; r < rounds; r++ {
		for _, pairing := range pairings {
			one := pairing.PlayerOne
			two := pairing.PlayerTwo
			oneName := pairing.PlayerOneName
			twoName := pairing.PlayerTwoName
			g := game.NewGame(one, two)
			wg.Add(1)
			go func() {
				defer wg.Done()
				outcome, err := g.Play(false)
				if err != nil {
					log.Print(err)
					return
				}
				deltaStatOne := PlayerStatistics{oneName, 1, 0, 0, 0, 0}
				deltaStatTwo := PlayerStatistics{twoName, 1, 0, 0, 0, 0}
				if outcome == board.PlayerOneWins {
					deltaStatOne.Won = 1
					deltaStatOne.Points = WinPoints
					deltaStatTwo.Lost = 1
				} else if outcome == board.PlayerTwoWins {
					deltaStatOne.Lost = 1
					deltaStatTwo.Won = 1
					deltaStatTwo.Points = WinPoints
				} else if outcome == board.Tie {
					deltaStatOne.Tied = 1
					deltaStatOne.Points = TiePoints
					deltaStatTwo.Tied = 1
					deltaStatTwo.Points = TiePoints
				}
				deltaStatChan <- deltaStatOne
				deltaStatChan <- deltaStatTwo
			}()
		}
	}
	go func() {
		wg.Wait()
		close(deltaStatChan)
	}()
	for deltaStat := range deltaStatChan {
		name := deltaStat.PlayerName
		if _, ok := stats[name]; !ok {
			log.Printf("no stats found for %s", deltaStat.PlayerName)
			continue
		}
		stats[name].Apply(&deltaStat)
	}
	result := make([]PlayerStatistics, 0)
	for _, stat := range stats {
		result = append(result, *stat)
	}
	return Result(result), nil
}

func pairUp(t *Tournament) []Pairing {
	pairings := make([]Pairing, 0)
	players := make([]Player, 0)
	for name, spawnFunc := range *t {
		players = append(players, Player{name, spawnFunc})
	}
	for i, leftPlayer := range players {
		for _, rightPlayer := range players[i+1:] {
			match := Pairing{
				leftPlayer.SpawnFunc(board.PlayerOne),
				leftPlayer.Name,
				rightPlayer.SpawnFunc(board.PlayerTwo),
				rightPlayer.Name,
			}
			rematch := Pairing{
				rightPlayer.SpawnFunc(board.PlayerOne),
				rightPlayer.Name,
				leftPlayer.SpawnFunc(board.PlayerTwo),
				leftPlayer.Name,
			}
			pairings = append(pairings, match)
			pairings = append(pairings, rematch)
		}
	}
	return pairings
}
