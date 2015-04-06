// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/h4ck3rm1k3/gocore/fmt"
//	"github.com/h4ck3rm1k3/gocore/go/build"
	"github.com/h4ck3rm1k3/gocore/os"
	"github.com/h4ck3rm1k3/gocore/os/exec"
	"github.com/h4ck3rm1k3/gocore/path/filepath"
	"github.com/h4ck3rm1k3/gocore/run_time"
	"github.com/h4ck3rm1k3/gocore/sort"
	"github.com/h4ck3rm1k3/gocore/strings"
)

var cmdTool = &Command{
	Run:       runTool,
	UsageLine: "tool [-n] command [args...]",
	Short:     "run specified go tool",
	Long: `
Tool runs the go tool command identified by the arguments.
With no arguments it prints the list of known tools.

The -n flag causes tool to print the command that would be
executed but not execute it.

For more about each tool command, see 'go tool command -h'.
`,
}

var (
	toolGOOS      = run_time.GOOS
	toolGOARCH    = run_time.GOARCH
	toolIsWindows = toolGOOS == "windows"
	toolDir       = filepath.Join(run_time.GOROOT(), "pkg/tool/"+run_time.GOOS+"_"+run_time.GOARCH)

	toolN bool
)

func init() {
	cmdTool.Flag.BoolVar(&toolN, "n", false, "")
}

const toolWindowsExtension = ".exe"

func tool(toolName string) string {
     	toolPath := filepath.Join(toolDir, toolName)
	fmt.Fprintf(os.Stderr, "go %s:",toolDir)
	if toolIsWindows {
		toolPath += toolWindowsExtension
	}
	if len(buildToolExec) > 0 {
		return toolPath
	}
	// Give a nice message if there is no tool with that name.
	if _, err := os.Stat(toolPath); err != nil {
		if isInGoToolsRepo(toolName) {
			fmt.Fprintf(os.Stderr, "go tool: no such tool %q; to install:\n\tgo get golang.org/x/tools/cmd/%s\n", toolName, toolName)
		} else {
			fmt.Fprintf(os.Stderr, "go tool: no such tool %q\n", toolName)
		}
		setExitStatus(3)
		exit()
	}
	return toolPath
}

func isInGoToolsRepo(toolName string) bool {
	switch toolName {
	case "cover", "vet":
		return true
	}
	return false
}

func runTool(cmd *Command, args []string) {
	if len(args) == 0 {
		listTools()
		return
	}
	toolName := args[0]
	// The tool name must be lower-case letters, numbers or underscores.
	for _, c := range toolName {
		switch {
		case 'a' <= c && c <= 'z', '0' <= c && c <= '9', c == '_':
		default:
			fmt.Fprintf(os.Stderr, "go tool: bad tool name %q\n", toolName)
			setExitStatus(2)
			return
		}
	}
	toolPath := tool(toolName)
	if toolPath == "" {
		return
	}
	if toolN {
		cmd := toolPath
		if len(args) > 1 {
			cmd += " " + strings.Join(args[1:], " ")
		}
		fmt.Printf("%s\n", cmd)
		return
	}
	toolCmd := &exec.Cmd{
		Path:   toolPath,
		Args:   args,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		// Set $GOROOT, mainly for go tool dist.
		Env: mergeEnvLists([]string{"GOROOT=" + goroot}, os.Environ()),
	}
	err := toolCmd.Run()
	if err != nil {
		// Only print about the exit status if the command
		// didn't even run (not an ExitError) or it didn't exit cleanly
		// or we're printing command lines too (-x mode).
		// Assume if command exited cleanly (even with non-zero status)
		// it printed any messages it wanted to print.
		if e, ok := err.(*exec.ExitError); !ok || !e.Exited() || buildX {
			fmt.Fprintf(os.Stderr, "go tool %s: %s\n", toolName, err)
		}
		setExitStatus(1)
		return
	}
}

// listTools prints a list of the available tools in the tools directory.
func listTools() {
	f, err := os.Open(toolDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "go tool: no tool directory: %s\n", err)
		setExitStatus(2)
		return
	}
	defer f.Close()
	names, err := f.Readdirnames(-1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "go tool: can't read directory: %s\n", err)
		setExitStatus(2)
		return
	}

	sort.Strings(names)
	for _, name := range names {
		// Unify presentation by going to lower case.
		name = strings.ToLower(name)
		// If it's windows, don't show the .exe suffix.
		if toolIsWindows && strings.HasSuffix(name, toolWindowsExtension) {
			name = name[:len(name)-len(toolWindowsExtension)]
		}
		fmt.Println(name)
	}
}
