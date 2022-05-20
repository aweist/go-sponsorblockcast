package dns

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/aweist/go-sponsorblockcast/models"
	"github.com/grandcat/zeroconf"
)

func Browse() models.ChromecastMapIface {
	chromecastMap := models.NewChromecastMap()
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			// printServiceEntry(entry)
			txt := textRecordMap(entry.Text)
			chromecastMap.Store(models.CastEntry{
				ID:   entry.Instance,
				Name: txt["fn"],
				Type: txt["md"],
				Port: entry.Port,
				IPv4: entry.AddrIPv4[0],
			})
		}
		log.Println("No more entries.")
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	err = resolver.Browse(ctx, "_googlecast._tcp", "local", entries)
	if err != nil {
		log.Fatalln("Failed to browse:", err.Error())
	}

	<-ctx.Done()
	return chromecastMap
}

func textRecordMap(records []string) map[string]string {
	m := map[string]string{}
	for _, s := range records {
		kv := strings.Split(s, "=")
		m[kv[0]] = kv[1]
	}
	return m
}

func printServiceEntry(se *zeroconf.ServiceEntry) {
	b, err := json.MarshalIndent(se, "", "  ")
	if err != nil {
		log.Println(err)
	}
	log.Println("Addrv4:", se.AddrIPv4)
	log.Print(string(b))
}
