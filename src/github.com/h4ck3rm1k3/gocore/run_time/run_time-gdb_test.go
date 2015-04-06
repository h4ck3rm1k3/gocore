package run_time_test

import (
	"github.com/h4ck3rm1k3/gocore/bytes"
	"github.com/h4ck3rm1k3/gocore/fmt"
	"github.com/h4ck3rm1k3/gocore/io/ioutil"
	"github.com/h4ck3rm1k3/gocore/os"
	"github.com/h4ck3rm1k3/gocore/os/exec"
	"github.com/h4ck3rm1k3/gocore/path/filepath"
	"github.com/h4ck3rm1k3/gocore/regexp"
	"github.com/h4ck3rm1k3/gocore/run_time"
	"testing"
)

func checkGdbPython(t *testing.T) {
	cmd := exec.Command("gdb", "-nx", "-q", "--batch", "-iex", "python import sys; print('go gdb python support')")
	out, err := cmd.CombinedOutput()

	if err != nil {
		t.Skipf("skipping due to issue running gdb: %v", err)
	}
	if string(out) != "go gdb python support\n" {
		t.Skipf("skipping due to lack of python gdb support: %s", out)
	}
}

const helloSource = `
package main
import "github.com/h4ck3rm1k3/gocore/fmt"
func main() {
	mapvar := make(map[string]string,5)
	mapvar["abc"] = "def"
	mapvar["ghi"] = "jkl"
	fmt.Println("hi") // line 8
}
`

func TestGdbPython(t *testing.T) {
	if run_time.GOOS == "darwin" {
		t.Skip("gdb does not work on darwin")
	}

	checkGdbPython(t)

	dir, err := ioutil.TempDir("", "go-build")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "main.go")
	err = ioutil.WriteFile(src, []byte(helloSource), 0644)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	cmd := exec.Command("go", "build", "-o", "a.exe")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("building source %v\n%s", err, out)
	}

	got, _ := exec.Command("gdb", "-nx", "-q", "--batch", "-iex",
		fmt.Sprintf("add-auto-load-safe-path %s/src/run_time", run_time.GOROOT()),
		"-ex", "br main.go:8",
		"-ex", "run",
		"-ex", "echo BEGIN info goroutines\n",
		"-ex", "info goroutines",
		"-ex", "echo END\n",
		"-ex", "echo BEGIN print mapvar\n",
		"-ex", "print mapvar",
		"-ex", "echo END\n",
		filepath.Join(dir, "a.exe")).CombinedOutput()

	firstLine := bytes.SplitN(got, []byte("\n"), 2)[0]
	if string(firstLine) != "Loading Go Runtime support." {
		t.Fatalf("failed to load Go run_time support: %s", firstLine)
	}

	// Extract named BEGIN...END blocks from output
	partRe := regexp.MustCompile(`(?ms)^BEGIN ([^\n]*)\n(.*?)\nEND`)
	blocks := map[string]string{}
	for _, subs := range partRe.FindAllSubmatch(got, -1) {
		blocks[string(subs[1])] = string(subs[2])
	}

	infoGoroutinesRe := regexp.MustCompile(`\*\s+\d+\s+running\s+`)
	if bl := blocks["info goroutines"]; !infoGoroutinesRe.MatchString(bl) {
		t.Fatalf("info goroutines failed: %s", bl)
	}

	printMapvarRe := regexp.MustCompile(`\Q = map[string]string = {["abc"] = "def", ["ghi"] = "jkl"}\E$`)
	if bl := blocks["print mapvar"]; !printMapvarRe.MatchString(bl) {
		t.Fatalf("print mapvar failed: %s", bl)
	}
}
