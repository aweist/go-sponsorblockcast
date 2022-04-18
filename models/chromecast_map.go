package models

import (
	"fmt"
	"sync"
)

type ChromecastMapIface interface {
	Store(CastEntry)
	Load(string) (CastEntry, error)
	Delete(string)
	Entries() map[string]CastEntry
}

type chromecastMap struct {
	m sync.Map
}

func NewChromecastMap() *chromecastMap {
	return &chromecastMap{
		m: sync.Map{},
	}
}

func (cm *chromecastMap) Store(castEntry CastEntry) {
	cm.m.Store(castEntry.ID, castEntry)
}

func (cm *chromecastMap) Load(id string) (CastEntry, error) {
	ce, ok := cm.m.Load(id)
	if !ok {
		return CastEntry{}, fmt.Errorf("no entry with id %s found", id)
	}
	castEntry, ok := ce.(CastEntry)
	if !ok {
		return CastEntry{}, fmt.Errorf("error asserting type of entry to CastEntry for id %s", id)
	}
	return castEntry, nil

}

func (cm *chromecastMap) Delete(id string) {
	cm.m.Delete(id)
}

func (cm *chromecastMap) Entries() map[string]CastEntry {
	result := map[string]CastEntry{}
	cm.m.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(CastEntry)
		result[k] = v
		return true
	})
	return result
}
