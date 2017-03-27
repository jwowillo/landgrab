// Package main runs a CLI landgrab interface.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/landgrab/player"
)

// shouldNotWait being true means the CLI doesn't ask for enter to be pressed to
// cintinue.
var shouldNotWait bool

const (
	// pieceCount for each game.Player.
	pieceCount = 5
	// damage for each game.Piece.
	damage = 1
	// life of each game.Piece.
	life = 3
	// damageIncrease applied whenever a game.Piece destroys an enemy
	// game.Piece.
	damageIncrease = 1
	// lifeIncrease applied whenever a game.Piece destroys an enemy
	// game.Piece.
	lifeIncrease = 1
)

// players available for use.
var players = map[string]game.Player{
	"greedy": player.NewGreedy(),
	"random": player.NewRandom(),
}

// main prompts for a game.Player 1 and 2 then plays a game of landgrab with the
// them.
func main() {
	w := bufio.NewWriter(os.Stdout)
	r := game.NewRules(
		pieceCount,
		damage, life,
		damageIncrease, lifeIncrease,
	)
	p1, p2 := choosePlayers(w)
	s := game.NewState(r, p1, p2)
	for s.Winner() == game.NoPlayer {
		printStateAndPrompt(w, s)
		w.Flush()
		if !shouldNotWait {
			waitForEnter(w)
		}
		s = game.NextState(s)
	}
	printState(w, s)
	w.Flush()
}

// init parses command-line flags.
func init() {
	flag.BoolVar(
		&shouldNotWait,
		"dont-wait",
		false,
		"doesn't wait for enter if true",
	)
	flag.Parse()
}

// choosePlayers prompts the user for two game.Players and returns the user's
// choice.
func choosePlayers(w *bufio.Writer) (game.Player, game.Player) {
	fmt.Fprintf(w, clear())
	fmt.Fprintln(w, title())
	p1 := choosePlayer(w, game.Player1)
	fmt.Fprintln(w)
	p2 := choosePlayer(w, game.Player2)
	return p1, p2
}

// choosePlayer prompts for a single game.Player for the game.PlayerID and
// returns the choice.
func choosePlayer(w *bufio.Writer, id game.PlayerID) game.Player {
	var p game.Player
	var playerNames []string
	for option := range players {
		playerNames = append(playerNames, option)
	}
	sort.Strings(playerNames)
	for p == nil {
		fmt.Fprintf(
			w,
			"Choose a player %s.\n",
			colorForPlayer(id)("%s", id),
		)
		for _, option := range playerNames {
			fmt.Fprintf(w, "* %s\n", option)
		}
		w.Flush()
		var choice string
		fmt.Scanf("%s", &choice)
		p = players[choice]
	}
	return p
}

// title string.
func title() string {
	return `
██╗      █████╗ ███╗   ██╗██████╗  ██████╗ ██████╗  █████╗ ██████╗
██║     ██╔══██╗████╗  ██║██╔══██╗██╔════╝ ██╔══██╗██╔══██╗██╔══██╗
██║     ███████║██╔██╗ ██║██║  ██║██║  ███╗██████╔╝███████║██████╔╝
██║     ██╔══██║██║╚██╗██║██║  ██║██║   ██║██╔══██╗██╔══██║██╔══██╗
███████╗██║  ██║██║ ╚████║██████╔╝╚██████╔╝██║  ██║██║  ██║██████╔╝
╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝
	`
}

// clear string.
func clear() string {
	return "\033[H\033[2J"
}

// board string.
func board(s *game.State) string {
	out := ""
	for i := 0; i < s.Rules().BoardSize(); i++ {
		for j := 0; j < s.Rules().BoardSize(); j++ {
			p := s.PieceForCell(game.NewCell(i, j))
			if p == game.NoPiece {
				out += "▒▒▒"
			} else {
				out += colorForPlayer(s.PlayerForPiece(p))(
					"%d%d%d",
					p.Damage(), p.ID(), p.Life(),
				)
			}
		}
		out += "\n"
	}
	return strings.TrimSpace(out)
}

// legend string.
func legend() string {
	out := ""
	for _, id := range []game.PlayerID{game.Player1, game.Player2} {
		out += fmt.Sprintf("%s: %s\n", id, colorForPlayer(id)("▒"))
	}
	out += "cell: DAMAGE|PIECE ID|LIFE"
	return out
}

// prompt string.
func prompt(s *game.State) string {
	return fmt.Sprintf("Current player: %s", s.CurrentPlayer())
}

// colorForPlayer returns a formatting function which formats a message and
// makes its result the appropriate color for the game.PlayerID.
func colorForPlayer(id game.PlayerID) func(string, ...interface{}) string {
	if id == game.Player1 {
		return red
	}
	if id == game.Player2 {
		return blue
	}
	return fmt.Sprintf
}

// red formatting function.
func red(s string, args ...interface{}) string {
	return fmt.Sprintf("\x1b[31;1m%s\x1b[0m", fmt.Sprintf(s, args...))
}

// blue formatting function.
func blue(s string, args ...interface{}) string {
	return fmt.Sprintf("\x1b[34;1m%s\x1b[0m", fmt.Sprintf(s, args...))
}

// printStateandPrompt prints the game.State and prompts to continue.
func printStateAndPrompt(w io.Writer, s *game.State) {
	printState(w, s)
	fmt.Fprintln(w)
	printPrompt(w, s)
}

// printState prints the game.State.
func printState(w io.Writer, s *game.State) {
	fmt.Fprint(w, clear())
	fmt.Fprintln(w, title())
	fmt.Fprintln(w, board(s))
	fmt.Fprintln(w)
	fmt.Fprintln(w, legend())
}

// printPrompt prints a prompt for the current game.Player.
func printPrompt(w io.Writer, s *game.State) {
	fmt.Fprintln(w, prompt(s))
}

// waitForEnter blocks until enter is pressed.
func waitForEnter(w *bufio.Writer) {
	fmt.Fprintln(w, "Press enter for next turn.")
	w.Flush()
	fmt.Scanf("%s")
}
