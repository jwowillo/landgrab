// Package cli ...
package cli

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/jwowillo/landgrab/convert"
	"github.com/jwowillo/landgrab/game"
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
func (cli *CLI) Run(factory *game.PlayerFactory, p1, p2 game.DescribedPlayer) {
	fmt.Fprintf(cli.rw, clear)
	fmt.Fprintf(cli.rw, title)
	fmt.Fprintln(cli.rw)
	cli.writeFunc()
	if p1 == nil {
		p1 = cli.choosePlayer(factory, game.Player1)
	}
	if p2 == nil {
		p2 = cli.choosePlayer(factory, game.Player2)
	}
	s := game.NewState(game.StandardRules, p1, p2)
	for s.Winner() == game.NoPlayer {
		printStateAndPrompt(cli.rw, s)
		cli.writeFunc()
		isHuman := false
		if s.CurrentPlayer() == game.Player1 && p1.Name() == "human" {
			play := cli.promptPlay(s)
			p1 = factory.SpecialPlayer("human", play)
			isHuman = true
		}
		if s.CurrentPlayer() == game.Player2 && p2.Name() == "human" {
			play := cli.promptPlay(s)
			p2 = factory.SpecialPlayer("human", play)
			isHuman = true
		}
		s = game.NewStateFromInfo(
			s.Rules(),
			s.CurrentPlayer(),
			p1, p2,
			s.Player1Pieces(), s.Player2Pieces(),
			pieces(s),
		)
		if cli.shouldWait && !isHuman {
			cli.waitForEnter()
		}
		s = game.NextState(s)
	}
	printState(cli.rw, s)
	cli.writeFunc()
}

func pieces(s *game.State) map[game.Cell]game.Piece {
	ps := make(map[game.Cell]game.Piece)
	for i := 0; i < s.Rules().BoardSize(); i++ {
		for j := 0; j < s.Rules().BoardSize(); j++ {
			c := game.NewCell(i, j)
			p := s.PieceForCell(c)
			if p == game.NoPiece {
				continue
			}
			ps[c] = p
		}
	}
	return ps
}

func (cli *CLI) promptPlay(s *game.State) map[string]interface{} {
	ms := make([][]game.Direction, s.Rules().PieceCount()*2)
	for _, m := range game.LegalMoves(s) {
		ms[m.Piece().ID()] = append(ms[m.Piece().ID()], m.Direction())
	}
	fmt.Fprintf(cli.rw, "\nLegal moves for play:\n")
	for i, p := range ms {
		if len(p) == 0 {
			continue
		}
		fmt.Fprintf(cli.rw, "* Piece %d: %v\n", i, p)
	}
	fmt.Fprintf(cli.rw, "\nEnter play as semi-colon separated pairs of piece ID and\n")
	fmt.Fprintf(cli.rw, "direction [(<id>,<direction>),(<id>,<direction>)...]:\n")
	cli.writeFunc()
	playString := ""
	fmt.Fscanf(cli.rw, "%s", &playString)
	pairs := strings.Split(playString, ";")
	play := make([]convert.JSONMove, len(pairs))
	for i, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if !strings.HasPrefix(pair, "(") || !strings.HasSuffix(pair, ")") {
			fmt.Fprintf(cli.rw, "\nInvalid play format.\n")
			cli.waitForEnter()
			return nil
		}
		pair = pair[1 : len(pair)-1]
		move := strings.Split(pair, ",")
		if len(move) != 2 {
			fmt.Fprintf(cli.rw, "\nInvalid play format.\n")
			cli.waitForEnter()
			return nil
		}
		move[0] = strings.TrimSpace(move[0])
		move[1] = strings.TrimSpace(move[1])
		id, err := strconv.Atoi(move[0])
		if err != nil {
			fmt.Fprintf(cli.rw, "\nInvalid play format.\n")
			cli.waitForEnter()
			return nil
		}
		var direction game.Direction
		switch move[1] {
		case "north":
			direction = game.North
		case "north-east":
			direction = game.NorthEast
		case "east":
			direction = game.East
		case "south-east":
			direction = game.SouthEast
		case "south":
			direction = game.South
		case "south-west":
			direction = game.SouthWest
		case "west":
			direction = game.West
		case "north-west":
			direction = game.NorthWest
		default:
			fmt.Fprintf(cli.rw, "\nInvalid play format.\n")
			cli.waitForEnter()
			return nil
		}
		piece := game.NoPiece
		for _, p := range s.CurrentPlayerPieces() {
			if int(p.ID()) == id {
				piece = p
			}
		}
		if piece == game.NoPiece {
			fmt.Fprintf(cli.rw, "\nInvalid play format.\n")
			cli.waitForEnter()
			return nil
		}
		play[i] = convert.MoveToJSONMove(
			game.NewMove(piece, direction),
			s,
		)
	}
	return map[string]interface{}{"moves": play}
}

// choosePlayer prompts for a single game.Player for the game.PlayerID and
// returns the choice.
func (cli *CLI) choosePlayer(
	factory *game.PlayerFactory, id game.PlayerID,
) game.DescribedPlayer {
	var p game.DescribedPlayer
	for p == nil {
		fmt.Fprintf(
			cli.rw,
			"\nChoose a player %s:\n",
			colorForPlayer(id)("%s", id),
		)
		for _, option := range factory.All() {
			fmt.Fprintf(cli.rw, "* %s\n", option)
		}
		cli.writeFunc()
		var choice string
		fmt.Fscanf(cli.rw, "%s", &choice)
		data := make(map[string]interface{})
		switch choice {
		case "api":
			fmt.Fprintf(cli.rw, "Enter URL: ")
			cli.writeFunc()
			var url string
			fmt.Fscanf(cli.rw, "%s", &url)
			data["url"] = url
		}
		p = factory.SpecialPlayer(choice, data)

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
				out += "▒▒▒▒▒▒"
			} else {
				out += colorForPlayer(s.PlayerForPiece(p))(
					"%2d|%d|%d",
					p.ID(), p.Life(), p.Damage(),
				)
			}
		}
		out += "\n"
	}
	return strings.TrimSpace(out)
}

// legend string.
func legend() string {
	return "cell: PIECE_ID|LIFE|DAMAGE"
}

// prompt string.
func prompt(s *game.State) string {
	return fmt.Sprintf("Current player: %s", colorForPlayer(s.CurrentPlayer())(s.CurrentPlayer().String()))
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
