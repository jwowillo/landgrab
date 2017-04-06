// Package cli ...
package cli

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/landgrab/player"
)

// CLI ...
type CLI struct {
	rw         io.ReadWriter
	writeFunc  func()
	shouldWait bool
}

// New ...
func New(r io.Reader, w io.Writer, wf func(), sw bool) *CLI {
	return &CLI{
		rw: struct {
			io.Reader
			io.Writer
		}{Reader: r, Writer: w},
		writeFunc:  wf,
		shouldWait: sw,
	}
}

// Run ...
//
// If player 1 or player 2 are nil, ask for them.
func (cli *CLI) Run(ps []player.Described, p1, p2 game.Player) {
	fmt.Fprintf(cli.rw, clear)
	fmt.Fprintf(cli.rw, title)
	fmt.Fprintln(cli.rw)
	cli.writeFunc()
	if p1 == nil {
		p1 = cli.choosePlayer(ps, game.Player1)
	}
	if p2 == nil {
		p2 = cli.choosePlayer(ps, game.Player2)
	}
	s := game.NewState(game.StandardRules, p1, p2)
	for s.Winner() == game.NoPlayer {
		printStateAndPrompt(cli.rw, s)
		cli.writeFunc()
		if cli.shouldWait {
			cli.waitForEnter()
		}
		s = game.NextState(s)
	}
	printState(cli.rw, s)
	cli.writeFunc()
}

// choosePlayer prompts for a single game.Player for the game.PlayerID and
// returns the choice.
func (cli *CLI) choosePlayer(
	ps []player.Described,
	id game.PlayerID,
) game.Player {
	players := make(map[string]game.Player)
	var p game.Player
	var playerNames []string
	for _, p := range ps {
		playerNames = append(
			playerNames,
			fmt.Sprintf("%s: %s", p.Name(), p.Description()),
		)
		players[p.Name()] = p
	}
	sort.Strings(playerNames)
	for p == nil {
		fmt.Fprintf(
			cli.rw,
			"Choose a player %s.\n",
			colorForPlayer(id)("%s", id),
		)
		for _, option := range playerNames {
			fmt.Fprintf(cli.rw, "* %s\n", option)
		}
		cli.writeFunc()
		var choice string
		fmt.Fscanf(cli.rw, "%s", &choice)
		p = players[choice]
	}
	return p
}

// waitForEnter blocks until enter is pressed.
func (cli *CLI) waitForEnter() {
	fmt.Fprintln(cli.rw)
	fmt.Fprintln(cli.rw, "Press enter for next turn.")
	cli.writeFunc()
	fmt.Fscanf(cli.rw, "%s")
}

// title string.
const title = `
██╗      █████╗ ███╗   ██╗██████╗  ██████╗ ██████╗  █████╗ ██████╗
██║     ██╔══██╗████╗  ██║██╔══██╗██╔════╝ ██╔══██╗██╔══██╗██╔══██╗
██║     ███████║██╔██╗ ██║██║  ██║██║  ███╗██████╔╝███████║██████╔╝
██║     ██╔══██║██║╚██╗██║██║  ██║██║   ██║██╔══██╗██╔══██║██╔══██╗
███████╗██║  ██║██║ ╚████║██████╔╝╚██████╔╝██║  ██║██║  ██║██████╔╝
╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝
	`

// clear string.
const clear = "\033[H\033[2J"

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
	fmt.Fprint(w, clear)
	fmt.Fprintln(w, title)
	fmt.Fprintln(w, board(s))
	fmt.Fprintln(w)
	fmt.Fprintln(w, legend())
}

// printPrompt prints a prompt for the current game.Player.
func printPrompt(w io.ReadWriter, s *game.State) {
	fmt.Fprintln(w, prompt(s))
}
