package misc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Prompter struct {
	reader bufio.Reader
}

func (prompter *Prompter) Init() {
	prompter.reader = *bufio.NewReader(os.Stdin)
}

func (prompter *Prompter) PromptUserOptions(prompt string, options []string) int {
	if &prompter.reader == nil {
		prompter.reader = *bufio.NewReader(os.Stdin)
	}

	ret := 0
	for {
		fmt.Println(prompt)
		for index, option := range options {
			fmt.Printf("%d. %s\n", index, option)
		}
		fmt.Print("> ")
		answer, err := prompter.reader.ReadString('\n')
		fmt.Println()
		if err != nil {
			fmt.Println("Something went wrong while taking in that input. Try again.\n")
			continue
		}
		// trim newline and whitespace off
		answer = strings.TrimSpace(answer[:len(answer)-1])

		answer_val, err := strconv.ParseInt(answer, 10, 0)
		if err != nil {
			fmt.Println("Error parsing your answer. Try again.\n")
			continue
		}

		if answer_val < 0 || int(answer_val) >= len(options) {
			fmt.Printf("%s is not a valid answer. Try again.\n\n", answer)
			continue
		}

		ret = int(answer_val)
		break
	}

	return ret
}

func (prompter *Prompter) PromptUserFreeForm(prompt string) string {
	if &prompter.reader == nil {
		prompter.reader = *bufio.NewReader(os.Stdin)
	}

	ret := ""

	for {
		fmt.Println(prompt)

		fmt.Print("> ")
		ans, err := prompter.reader.ReadString('\n')
		fmt.Println()
		if err != nil {
			fmt.Println("Something went wrong collecting your answer. Try again.\n")
			continue
		}

		ans = strings.TrimSpace(ans[:len(ans)-1])

		if len(ans) == 0 {
			fmt.Println("You must provide an answer with at least one character of length\n")
			continue
		}

		ret = ans
		break
	}

	return ret
}
