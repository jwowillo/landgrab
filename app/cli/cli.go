// Package cli ...
package cli

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/jwowillo/landgrab/game"
)

// ReadWriter ...
type ReadWriter struct {
	io.Reader
	io.Writer
	writeFunc  func()
	shouldWait bool
}

// NewReadWriter ...
func NewReadWriter(r io.Reader, w io.Writer, wf func(), sw bool) *ReadWriter {
	return &ReadWriter{
		Reader:     r,
		Writer:     w,
		writeFunc:  wf,
		shouldWait: sw,
	}
}

// Run ...
//
// If player 1 or player 2 are nil, ask for them.
func Run(w *ReadWriter, ps map[string]game.Player, p1, p2 game.Player) {
	fmt.Fprintf(w, clear())
	fmt.Fprintf(w, title())
	fmt.Fprintln(w)
	w.writeFunc()
	if p1 == nil {
		p1 = choosePlayer(w, ps, game.Player1)
	}
	if p2 == nil {
		p2 = choosePlayer(w, ps, game.Player2)
	}
	s := game.NewState(game.StandardRules, p1, p2)
	for s.Winner() == game.NoPlayer {
		printStateAndPrompt(w, s)
		w.writeFunc()
		if w.shouldWait {
			waitForEnter(w)
		}
		s = game.NextState(s)
	}
	printState(w, s)
	w.writeFunc()
}

// choosePlayer prompts for a single game.Player for the game.PlayerID and
// returns the choice.
func choosePlayer(
	w *ReadWriter,
	ps map[string]game.Player,
	id game.PlayerID,
) game.Player {
	var p game.Player
	var playerNames []string
	for option := range ps {
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
		w.writeFunc()
		var choice string
		fmt.Fscanf(w, "%s", &choice)
		p = ps[choice]
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
func printStateAndPrompt(w io.ReadWriter, s *game.State) {
	printState(w, s)
	fmt.Fprintln(w)
	printPrompt(w, s)
}

// printState prints the game.State.
func printState(w io.ReadWriter, s *game.State) {
	fmt.Fprint(w, clear())
	fmt.Fprintln(w, title())
	fmt.Fprintln(w, board(s))
	fmt.Fprintln(w)
	fmt.Fprintln(w, legend())
}

// printPrompt prints a prompt for the current game.Player.
func printPrompt(w io.ReadWriter, s *game.State) {
	fmt.Fprintln(w, prompt(s))
}

// waitForEnter blocks until enter is pressed.
func waitForEnter(w *ReadWriter) {
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Press enter for next turn.")
	w.writeFunc()
	fmt.Fscanf(w, "%s")
}
