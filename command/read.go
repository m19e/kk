package command

import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/urfave/cli"
)

func CmdRead(c *cli.Context) error {
	// Write your code here
	
	// p, _ := os.Getwd()
	// fmt.Println(p)

	// ls, _ := ioutil.ReadDir("./")
	// for _, file := range ls {
	// 	// fmt.Println(file.Name())
	// }

	fmt.Println("read file...")
	// open file
	f, err := os.Open(c.Args().Get(0))
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	defer fmt.Println("")

	// read all text
	b, _ := ioutil.ReadAll(f)

	// output
	fmt.Println(string(b))

	return nil
}
