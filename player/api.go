package player

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/jwowillo/landgrab/convert"
	"github.com/jwowillo/landgrab/game"
)

// API game.DescribedPlayer consults an external API for game.Plays for each
// game.State.
//
// The API is expected to receive a query-string encoded convert.JSONState in
// JSON form and return a convert.JSONPlay in JSON form wrapped in a data
// payload.
//
// API is a special game.DescribedPlayer in that it needs its URL initialized.
type API struct {
	url    string
	client *http.Client
}

// newAPI creates an uninitialized game.DescribedPlayer.
func newAPI() game.DescribedPlayer {
	return &API{}
}

// SetURL of the API the game.DescribedPlayer consults.
func (p *API) SetURL(url string) {
	p.url = url
}

// Name returns "api".
func (p *API) Name() string {
	return "api"
}

// Description of the game.DescribedPlayer.
func (p *API) Description() string {
	return `passes API at URL a State at key "state" and makes the returned
	Play embedded in an object under key "data"`
}

// Play passes the game.State to an external API using the format described in
// the convert package and returns the game.Play the API returns.
//
// An empty game.Play is returned as nil if the external API doesnt return a
// game.Play correctly.
func (p *API) Play(s *game.State) game.Play {
	if p.client == nil {
		p.client = &http.Client{Transport: &http.Transport{
			Dial: makeTimeout(s.Rules().TimerDuration()),
		}}
	}
	js := convert.StateToJSONState(s)
	bs, err := json.Marshal(js)
	if err != nil {
		log.Println(err)
		return nil
	}
	query := "?state=" + url.QueryEscape(string(bs))
	resp, err := p.client.Get(p.url + query)
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
	raw := make(map[string]interface{})
	if err = json.Unmarshal(body, &raw); err != nil {
		log.Println(err)
		return nil
	}
	bs, err = json.Marshal(raw["data"])
	if err != nil {
		log.Println(err)
		return nil
	}
	play, err := convert.JSONToPlay(bs)
	if err != nil {
		log.Println(err)
		return nil
	}
	return play
}

// makeTimeout makes a net.Con which waits the given time.Duration before timing
// out.
func makeTimeout(d time.Duration) func(string, string) (net.Conn, error) {
	return func(network, addr string) (net.Conn, error) {
		return net.DialTimeout(network, addr, d)
	}
}
