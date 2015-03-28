package pkg_test

import "github.com/h4ck3rm1k3/gocore/os"

func init() {
	os.Stdout.Write([]byte("File with non-runnable example was built.\n"))
}

func Example_test() {
	// This test will not be run, it has no "Output:" comment.
}
