package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/jwowillo/landgrab/app/cli"
	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/landgrab/player"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	rw := cli.NewReadWriter(os.Stdin, w, func() { w.Flush() }, shouldWait)
	for _, key := range []string{player1, player2} {
		if _, ok := players[key]; key != "" && !ok {
			fmt.Fprintf(rw, "%s isn't a valid player\n", key)
			w.Flush()
			os.Exit(1)
		}
	}
	cli.Run(rw, players, players[player1], players[player2])
}

var (
	// shouldWait being true means the CLI doesn't ask for enter to be
	// pressed to continue.
	shouldWait       bool
	player1, player2 string
)

// players available for use.
var players = map[string]game.Player{
	"greedy": player.NewGreedy(),
	"random": player.NewRandom(),
}

// init parses command-line flags.
func init() {
	flag.BoolVar(
		&shouldWait,
		"wait",
		true,
		"waits for enter if true",
	)
	flag.StringVar(
		&player1,
		"player1",
		"",
		"choice for player1",
	)
	flag.StringVar(
		&player2,
		"player2",
		"",
		"choice for player2",
	)
	flag.Parse()
}
