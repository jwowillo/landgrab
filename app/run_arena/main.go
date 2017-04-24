package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jwowillo/landgrab/arena"
	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/landgrab/player"
)

func main() {
	p1 := buildPlayer(player1, player.Factory)
	p2 := buildPlayer(player2, player.Factory)
	if p1 == nil || p2 == nil {
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

func buildPlayer(name string, factory *game.PlayerFactory) game.DescribedPlayer {
	if name == "human" {
		return nil
	}
	data := make(map[string]interface{})
	if strings.HasPrefix(name, "api") {
		parts := strings.Split(name, "api:")
		if len(parts) != 2 {
			fmt.Println("invalid api format")
			os.Exit(1)
		}
		name = parts[0]
		data["url"] = parts[1]
	}
	return factory.SpecialPlayer(name, data)
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
