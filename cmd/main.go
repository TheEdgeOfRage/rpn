package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"gitea.theedgeofrage.com/theedgeofrage/rpn/rpn"
)

func main() {
	rpn := rpn.NewRPN()
	reader := bufio.NewReader(os.Stdin)
	for {
		rpn.PrintStack()
		// input := ""
		// _, err := fmt.Scanln(&input)

		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println()
				return
			}
			fmt.Println(err)
			continue
		}

		err = rpn.Eval(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
