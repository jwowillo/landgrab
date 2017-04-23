package game

// PlayerConstructor ...
type PlayerConstructor func() DescribedPlayer

// PlayerInitializer ...
type PlayerInitializer func(DescribedPlayer, map[string]interface{})

// PlayerFactory ...
type PlayerFactory struct {
	players      map[string]PlayerConstructor
	initializers map[string]PlayerInitializer
}

// NewPlayerFactory ...
func NewPlayerFactory() *PlayerFactory {
	return &PlayerFactory{
		players:      make(map[string]PlayerConstructor),
		initializers: make(map[string]PlayerInitializer),
	}
}

// All ...
func (f *PlayerFactory) All() []string {
	all := make([]string, len(f.players))
	i := 0
	for name := range f.players {
		all[i] = name
		i++
	}
	return all
}

// Player ...
func (f *PlayerFactory) Player(name string) DescribedPlayer {
	if _, ok := f.players[name]; !ok {
		return nil
	}
	return f.players[name]()
}

// SpecialPlayer ...
func (f *PlayerFactory) SpecialPlayer(
	name string,
	data map[string]interface{},
) DescribedPlayer {
	p := f.Player(name)
	if p == nil {
		return nil
	}
	if _, ok := f.initializers[name]; !ok {
		return p
	}
	f.initializers[name](p, data)
	return p
}

// Register ...
func (f *PlayerFactory) Register(ctor PlayerConstructor) {
	f.players[ctor().Name()] = ctor
}

// RegisterSpecial ...
func (f *PlayerFactory) RegisterSpecial(
	ctor PlayerConstructor,
	init PlayerInitializer,
) {
	f.Register(ctor)
	f.initializers[ctor().Name()] = init
}

// DescribedPlayer ...
type DescribedPlayer interface {
	Player
	Name() string
	Description() string
}

// Player in the game.
//
// The Player will be asked to return their choice of Play for each State where
// it is their turn in a game.
type Player interface {
	Play(*State) Play
}

// PlayerID identifies a Player in a game.
type PlayerID int

// PlayerIDs which can be given to Players.
const (
	NoPlayer PlayerID = iota
	Player1
	Player2
)

// String representation of the PlayerID.
func (id PlayerID) String() string {
	return map[PlayerID]string{
		NoPlayer: "no player",
		Player1:  "player 1",
		Player2:  "player 2",
	}[id]
}
