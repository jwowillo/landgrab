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
	app := cli.New(os.Stdin, w, func() { w.Flush() }, shouldWait)
	players := make(map[string]game.Player)
	for _, p := range player.All() {
		players[p.Name()] = p
	}
	for _, key := range []string{player1, player2} {
		if _, ok := players[key]; key != "" && !ok {
			fmt.Fprintf(w, "%s isn't a valid player\n", key)
			w.Flush()
			os.Exit(1)
		}
	}
	app.Run(player.All(), players[player1], players[player2])
}

var (
	// shouldWait being true means the CLI doesn't ask for enter to be
	// pressed to continue.
	shouldWait       bool
	player1, player2 string
)

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
