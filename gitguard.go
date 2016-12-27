// Original template https://gist.github.com/MakoTano/624fe3fdea914b262e2c
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gitguard"
	app.Usage = "github.com/akm/gitguard"
	app.Version = "0.0.1"

	app.Action = executeCommand

	app.Commands = []cli.Command{
		{
			Name: "status",
			Aliases: []string{"s"},
			Usage: "Show status of this repository",
			Action: showStatus,
		},
	}

	app.Run(os.Args)
}

func executeCommand(c *cli.Context) error {

	if c.NArg() < 1 {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	if !isClean() {
		fmt.Println("\x1b[31m[gitguard] There are files that need to be committed first.\x1b[0m")
		fmt.Println("[gitguard] git status")
		runCommand("git", "status")
		os.Exit(1)
	}

	var args []string = c.Args()
	runCommandWithExit(args...)

	runCommandWithExit("git", "add", ".")
	runCommandWithExit("git", "commit", "-m", strings.Join(args, " "))

	return nil
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func runCommandWithExit(args ...string) {
	var cmdStr = strings.Join(args, " ")
	fmt.Println("[gitguard] " + cmdStr)
	err := runCommand(args[0], args[1:]...)
	if err != nil {
		fmt.Printf("\x1b[31m[gitguard]%v\x1b[0m\n", err)
		os.Exit(1)
	}
}

func showStatus(c *cli.Context) error {
	if isClean() {
		fmt.Printf("\x1b[32mClean\x1b[0m\n")
	} else {
		fmt.Printf("\x1b[31mNG\x1b[0m\n")
	}
  return nil
}

func isClean() bool {
  return !isChanged() && noUntrackedFiles()
}

func isChanged() bool {
	cmd := exec.Command("git", "diff", "--exit-code")
	cmd.Stdout = nil
	err := cmd.Run()
	return err != nil
}

func noUntrackedFiles() bool {
	cmd := exec.Command("git", "ls-files", " --others", "--exclude-standard")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return len(out) == 0
}
