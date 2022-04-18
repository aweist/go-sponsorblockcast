package models

import "net"

type CastEntry struct {
	ID   string
	Name string
	Type string
	Port int
	IPv4 net.IP
}
