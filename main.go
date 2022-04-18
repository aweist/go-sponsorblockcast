package main

import (
	"github.com/aweist/go-sponsorblockcast/dns"
	"github.com/aweist/go-sponsorblockcast/models"
	"github.com/aweist/go-sponsorblockcast/server"
)

func main() {
	chromecastMap := models.NewChromecastMap()
	go dns.Browse(chromecastMap)
	server.Serve(chromecastMap)
}
