package command

import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/urfave/cli"
	"regexp"
)

func isKanji(r rune) bool {
	reg := regexp.MustCompile(`[\x{3400}-\x{9FFF}]`)
	return reg.MatchString(string(r))
}

func CmdRead(c *cli.Context) error {
	// Write your code here
	
	// p, _ := os.Getwd()
	// fmt.Println(p)

	// ls, _ := ioutil.ReadDir("./")
	// for _, file := range ls {
	// 	// fmt.Println(file.Name())
	// }

	fmt.Println(">> read file...")
	// open file
	f, err := os.Open(c.Args().Get(0))
	if err != nil {
		fmt.Println(err)
	}

	defer fmt.Println(">> finish")
	defer f.Close()

	// read all text
	b, _ := ioutil.ReadAll(f)

	// output
	// fmt.Println(string(b))

	for _, char := range string(b) {
		if isKanji(char) {
			fmt.Printf("%s", string(char))
		} else {
			//fmt.Printf("%s", string(char))
		}
	}

	fmt.Println("")
	return nil
}
