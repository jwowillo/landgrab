package arena

import "github.com/jwowillo/landgrab/game"

// CumulativeResult ...
type CumulativeResult struct {
	Player1Wins          int
	Player2Wins          int
	Player1AveragePieces float64
	Player1AverageLife   float64
	Player1AverageDamage float64
	Player2AveragePieces float64
	Player2AverageLife   float64
	Player2AverageDamage float64
	AverageTurns         float64
}

// Result ....
type Result struct {
	Winner        game.PlayerID
	Player1Pieces float64
	Player1Life   float64
	Player1Damage float64
	Player2Pieces float64
	Player2Life   float64
	Player2Damage float64
	Turns         int
}

// Run ...
func Run(rules game.Rules, p1, p2 game.Player, n int) CumulativeResult {
	results := make(chan Result)
	for i := 0; i < n; i++ {
		go func(results chan Result) {
			results <- RunSingle(rules, p1, p2)
		}(results)
	}
	result := CumulativeResult{}
	for i := 0; i < n; i++ {
		r := <-results
		if r.Winner == game.Player1 {
			result.Player1Wins++
		}
		if r.Winner == game.Player2 {
			result.Player2Wins++
		}
		result.Player1AveragePieces += r.Player1Pieces
		result.Player1AverageLife += r.Player1Life
		result.Player1AverageDamage += r.Player1Damage
		result.Player2AveragePieces += r.Player2Pieces
		result.Player2AverageLife += r.Player2Life
		result.Player2AverageDamage += r.Player2Damage
		result.AverageTurns += float64(r.Turns)
	}
	result.Player1AveragePieces /= float64(n)
	result.Player1AverageLife /= float64(n)
	result.Player1AverageDamage /= float64(n)
	result.Player2AveragePieces /= float64(n)
	result.Player2AverageLife /= float64(n)
	result.Player2AverageDamage /= float64(n)
	result.AverageTurns /= float64(n)
	return result
}

// RunSingle ...
func RunSingle(rules game.Rules, p1, p2 game.Player) Result {
	r := Result{}
	s := game.NewState(rules, p1, p2)
	for s.Winner() == game.NoPlayer {
		s = game.NextState(s)
		r.Turns++
	}
	r.Winner = s.Winner()
	for _, p := range s.Player1Pieces() {
		r.Player1Pieces++
		r.Player1Life += float64(p.Life())
		r.Player1Damage += float64(p.Damage())
	}
	for _, p := range s.Player2Pieces() {
		r.Player2Pieces++
		r.Player2Life += float64(p.Life())
		r.Player2Damage += float64(p.Damage())
	}
	if len(s.Player1Pieces()) != 0 {
		r.Player1Life /= float64(len(s.Player1Pieces()))
		r.Player1Damage /= float64(len(s.Player1Pieces()))
	}
	if len(s.Player2Pieces()) != 0 {
		r.Player2Life /= float64(len(s.Player2Pieces()))
		r.Player2Damage /= float64(len(s.Player2Pieces()))
	}
	return r

}
