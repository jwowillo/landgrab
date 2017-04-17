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
	players := make(map[string]game.DescribedPlayer)
	for _, p := range player.All() {
		players[p.Name()] = p
	}
	trimmed1 := player1
	if strings.Contains(player1, ":") {
		trimmed1 = strings.Split(player1, ":")[0]
	}
	trimmed2 := player2
	if strings.Contains(player1, ":") {
		trimmed2 = strings.Split(player2, ":")[0]
	}
	p1, p1Ok := players[trimmed1]
	p2, p2Ok := players[trimmed2]
	if !p1Ok || !p2Ok {
		fmt.Fprintln(w, "invalid players chosen")
		w.Flush()
		os.Exit(1)
	}
	if strings.HasPrefix(player1, "api") {
		parts := strings.Split(player1, "api:")
		if len(parts) != 2 {
			fmt.Println("invalid api format")
			os.Exit(1)
		}
		url := parts[1]
		p1.(*player.API).SetURL(url)
	}
	if strings.HasPrefix(player2, "api") {
		parts := strings.Split(player2, "api:")
		if len(parts) != 2 {
			fmt.Println("invalid api format")
			os.Exit(1)
		}
		url := parts[1]
		p2.(*player.API).SetURL(url)
	}
	app.Run(player.All(), p1, p2)
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
