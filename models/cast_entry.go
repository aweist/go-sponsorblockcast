package models

import (
	"encoding/json"
	"log"
	"net"
)

type CastEntry struct {
	ID   string
	Name string
	Type string
	Port int
	IPv4 net.IP
}

func (ce CastEntry) Print() {
	json, err := json.MarshalIndent(ce, "", "  ")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(json))
}
