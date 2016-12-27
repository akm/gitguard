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

	app.Action = helloAction

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

func helloAction(c *cli.Context) error {

	// グローバルオプション
	var isDry = c.GlobalBool("dryrun")
	if isDry {
		fmt.Println("this is dry-run")
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
		fmt.Printf("Clean\n")
	} else {
		fmt.Printf("NG\n")
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
	return true
}
