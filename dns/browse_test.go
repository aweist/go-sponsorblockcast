package dns

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrowse(t *testing.T) {
	chromecastMap := Browse()
	assert.Greater(t, len(chromecastMap.Entries()), 0)

	for _, entry := range chromecastMap.Entries() {
		entry.Print()
	}
	fmt.Println()
}
