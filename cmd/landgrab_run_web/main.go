package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/jwowillo/landgrab/web"
	"github.com/jwowillo/trim/server"
)

func main() {
	s := server.New(url, port)
	s.AddHeader("Access-Control-Allow-Origin", "*")
	s.AddHeader("Access-Control-Allow-Headers", "Authorization")
	if port != 80 {
		url += fmt.Sprintf(":%d", port)
	}
	s.Serve(web.New("", url, filepath.Join("build", "web")))
}

var (
	url  string
	port int
)

func init() {
	flag.StringVar(&url, "url", "localhost", "URL to listen from")
	flag.IntVar(&port, "port", 5000, "port to run on")
	flag.Parse()
}
