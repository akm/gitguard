// Original template https://gist.github.com/MakoTano/624fe3fdea914b262e2c
package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "sample_client"
	app.Usage = "github.com/codegangsta/cli のテンプレートです"
	app.Version = "0.0.1"

	// グローバルオプション設定
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "dryrun, d", // 省略指定 => d
			Usage: "グローバルオプション dryrunです。",
		},
	}

	app.Action = executeCommand

	app.Commands = []cli.Command{
		{
			Name: "status",
			Aliases: []string{"s"},
			Usage: "Show status of this repository",
			Action: showStatus,
		},
	}

	app.Before = func(c *cli.Context) error {
		// 開始前の処理をここに書く
		fmt.Println("開始")
		return nil // error を返すと処理全体が終了
	}

	app.After = func(c *cli.Context) error {
		// 終了時の処理をここに書く
		fmt.Println("終了")
		return nil
	}

	app.Run(os.Args)
}

func executeCommand(c *cli.Context) error {

	if !isClean() {
		fmt.Println("\x1b[31m[gitguard] There are files that need to be committed first.\x1b[0m")
		fmt.Println("[gitguard] git status")
		cmd := exec.Command("git", "status")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		os.Exit(1)
	}

	// パラメータ
	var paramFirst = ""
	if len(c.Args()) > 0 {
		paramFirst = c.Args().First() // c.Args()[0] と同じ意味
	}

	fmt.Printf("Hello world! %s\n", paramFirst)
	return nil
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
