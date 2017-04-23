package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jwowillo/landgrab/cli"
	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/landgrab/player"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	app := cli.New(os.Stdin, w, func() { w.Flush() }, shouldWait)
	p1 := buildPlayer(w, player1, player.Factory)
	p2 := buildPlayer(w, player2, player.Factory)
	app.Run(player.Factory, p1, p2)
}

func buildPlayer(w *bufio.Writer, name string, factory *game.PlayerFactory) game.DescribedPlayer {
	data := make(map[string]interface{})
	if strings.HasPrefix(name, "api") {
		parts := strings.Split(name, "api:")
		if len(parts) != 2 {
			fmt.Fprintln(w, "invalid api format")
			w.Flush()
			os.Exit(1)
		}
		name = parts[0]
		data["url"] = parts[1]
	}
	return factory.SpecialPlayer(name, data)
}

var (
	// shouldWait being true means the CLI doesn't ask for enter to be
	// pressed to continue.
	shouldWait       bool
	player1, player2 string
)

// init parses command-line flags.
func init() {
	flag.BoolVar(&shouldWait, "wait", true, "waits for enter if true")
	flag.StringVar(&player1, "player1", "", "choice for player 1")
	flag.StringVar(&player2, "player2", "", "choice for player 2")
	flag.Parse()
}
