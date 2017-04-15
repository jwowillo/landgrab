package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jwowillo/landgrab/arena"
	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/landgrab/player"
)

func main() {
	if player1 == "human" || player2 == "human" {
		fmt.Println("human isn't a valid player")
		os.Exit(1)
	}
	players := make(map[string]game.Player)
	for _, p := range player.All() {
		if p.Name() == "human" {
			continue
		}
		players[p.Name()] = p
	}
	p1, p1Ok := players[player1]
	p2, p2Ok := players[player2]
	if !p1Ok || !p2Ok {
		fmt.Println("invalid players chosen")
		os.Exit(1)
	}
	if n < 0 {
		fmt.Println("n must be non-negative")
		os.Exit(1)
	}
	r := arena.Run(game.StandardRules, p1, p2, n)
	fmt.Println("Player 1 Wins:", r.Player1Wins)
	fmt.Println("Player 1 Average Pieces:", r.Player1AveragePieces)
	fmt.Println("Player 1 Average Life:", r.Player1AverageLife)
	fmt.Println("Player 1 Average Damage:", r.Player1AverageDamage)
	fmt.Println("Player 2 Wins:", r.Player2Wins)
	fmt.Println("Player 2 Average Pieces:", r.Player2AveragePieces)
	fmt.Println("Player 2 Average Life:", r.Player2AverageLife)
	fmt.Println("Player 2 Average Damage:", r.Player2AverageDamage)
	fmt.Println("Average Turns:", r.AverageTurns)
}

var (
	player1 string
	player2 string
	n       int
)

func init() {
	flag.StringVar(&player1, "player1", "", "choice for player 2")
	flag.StringVar(&player2, "player2", "", "choice for player 2")
	flag.IntVar(&n, "n", -1, "times to play")
	flag.Parse()
}
