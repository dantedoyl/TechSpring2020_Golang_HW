package main

import (
	"fmt"
	ap "unique/args_parser"
	rw "unique/read_write"
	"unique/uniq"
)

func main() {
	option, pareseErr := ap.ArgsParser()
	if pareseErr != nil {
		fmt.Println(pareseErr.Error())
	}

	text, readErr := rw.ReadText(option.InputType)
	if readErr != nil {
		fmt.Println(readErr.Error())
	}

	uniqStr := uniq.Uniq(text, *option)

	writeErr := rw.WriteText(uniqStr, option.OutputType)
	if writeErr != nil {
		fmt.Println(writeErr.Error())
	}
}
