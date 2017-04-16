package player

import (
	"reflect"
	"strings"

	"github.com/jwowillo/landgrab/game"
)

// API ...
type API struct {
	url string
}

// NewAPI ...
func NewAPI() API {
	return API{}
}

// SetURL ...
func (p API) SetURL(url string) {
	p.url = url
}

// Name ...
func (p API) Name() string {
	return strings.ToLower(reflect.TypeOf(p).Name())
}

// Description ...
func (p API) Description() string {
	return "asks API at the set URL for the play to make"
}

// Play ...
func (p API) Play(s *game.State) game.Play {
	return nil
}
