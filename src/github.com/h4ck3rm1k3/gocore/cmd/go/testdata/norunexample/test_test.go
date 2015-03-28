package pkg

import (
	"github.com/h4ck3rm1k3/gocore/os"
	"testing"
)

func TestBuilt(t *testing.T) {
	os.Stdout.Write([]byte("A normal test was executed.\n"))
}
