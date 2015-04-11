package main

import (
	"fmt"
	"github.com/h4ck3rm1k3/gocore/os"
)

func main() {
	fmt.Printf("test")
	for a, b := range os.Args  {
		fmt.Printf("a %v:%v\n",a, b)
	}
	
}
