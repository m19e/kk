package command

import (
	"fmt"
	"github.com/urfave/cli"
)

func CmdRead(c *cli.Context) error {
	// Write your code here
	fmt.Println("Hello friend!")
	return nil
}
