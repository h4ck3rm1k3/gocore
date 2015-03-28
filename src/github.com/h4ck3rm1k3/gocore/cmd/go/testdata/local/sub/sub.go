package sub

import (
	"github.com/h4ck3rm1k3/gocore/fmt"

	subsub "./sub"
)

func Hello() {
	fmt.Println("sub.Hello")
	subsub.Hello()
}
