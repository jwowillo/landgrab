package player

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/jwowillo/landgrab/convert"
	"github.com/jwowillo/landgrab/game"
)

// API ...
type API struct {
	url string
}

// NewAPI ...
func NewAPI() *API {
	return &API{}
}

// SetURL ...
func (p *API) SetURL(url string) {
	p.url = url
}

// Name ...
func (p *API) Name() string {
	return "api"
}

// Description ...
func (p *API) Description() string {
	return "asks API at the set URL for the play to make"
}

// Play ...
func (p *API) Play(s *game.State) game.Play {
	js := convert.StateToJSONState(s)
	bs, err := json.Marshal(js)
	if err != nil {
		log.Println(err)
		return nil
	}
	query := "?" + url.QueryEscape(string(bs))
	resp, err := http.Get(p.url + query)
	if err != nil {
		log.Println(err)
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	err = resp.Body.Close()
	if err != nil {
		log.Println(err)
		return nil
	}
	play, err := convert.JSONToPlay(body)
	if err != nil {
		log.Println(err)
		return nil
	}
	return play
}
