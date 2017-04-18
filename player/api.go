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

// API ...
type API struct {
	url    string
	client *http.Client
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
	return `passes API at URL a State at key "state" and makes the returned
	Play embedded in an object under key "data"`
}

// Play ...
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

func makeTimeout(d time.Duration) func(string, string) (net.Conn, error) {
	return func(network, addr string) (net.Conn, error) {
		return net.DialTimeout(network, addr, d)
	}
}
