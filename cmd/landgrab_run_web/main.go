package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/jwowillo/landgrab/web"
	"github.com/jwowillo/trim/server"
)

func main() {
	s := server.New(host, port)
	s.AddHeader("Access-Control-Allow-Origin", "*")
	s.AddHeader("Access-Control-Allow-Headers", "Authorization")
	if port != 80 {
		host += fmt.Sprintf(":%d", port)
	}
	s.Serve(web.New("", host, filepath.Join("build", "web")))
}

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "host", "localhost", "host to run on")
	flag.IntVar(&port, "port", 5000, "port to run on")
	flag.Parse()
}
