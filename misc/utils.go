package misc

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
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
	if prompter.reader.Size() == 0 {
		prompter.reader = *bufio.NewReader(os.Stdin)
	}

	ret := ""

	for {
		fmt.Printf("%s\n> ", prompt)
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

func (prompter *Prompter) PromptUserFloat(prompt string, currency bool) float64 {
	if &prompter.reader == nil {
		prompter.reader = *bufio.NewReader(os.Stdin)
	}

	var placeholder string
	if currency {
		placeholder = "> $"
	} else {
		placeholder = "> "
	}

	ret := float64(0)

	for {
		fmt.Printf("%s\n%s ", prompt, placeholder)

		ans, err := prompter.reader.ReadString('\n')
		fmt.Println()
		if err != nil {
			fmt.Println("There was an error reading your answer. Try again.\n")
			continue
		}

		ansValue, err := strconv.ParseFloat(strings.TrimSpace(ans[:len(ans)-1]), 64)
		if err != nil {
			fmt.Printf("%s could not be parsed as a float. Please provide a valid float.\n\n", ans)
			continue
		}

		ret = ansValue
		break
	}

	return ret
}

func (prompter *Prompter) PromptUserEnterKey(prompt string) {
	fmt.Println(prompt)
	fmt.Scanln()
}

func FormatCurrency(val any) (string, error) {
	typeOfVal := reflect.ValueOf(val)

	switch typeOfVal.Kind() {
	case reflect.String:
		floatVal, err := strconv.ParseFloat(val.(string), 64)
		if err != nil {
			fmt.Printf("There was an error formatting your currency of %s\n", val)
			return "", err
		}

		if floatVal < 0 {
			return fmt.Sprintf("-$%.2f", floatVal*-1), nil
		}
		return fmt.Sprintf("$%.2f", floatVal), nil
	case reflect.Float64, reflect.Float32:
		if val.(float64) < 0 {
			return fmt.Sprintf("-$%.2f", val.(float64)*-1), nil
		}

		return fmt.Sprintf("$%.2f", val), nil
	default:
		errorMsg := fmt.Sprintf("Currently unable to format a value of the type %s", typeOfVal.Kind())
		fmt.Println(errorMsg)
		return "", errors.New(errorMsg)
	}
}
