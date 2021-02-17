package args_parser

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type MutuallyExclusiveFlag struct {
	CountFlag bool
	DuplicateFlag bool
	UniqueFlag bool
}

type Option struct {
	MutuallyExcFlag MutuallyExclusiveFlag
	SkipWordFlag    int
	SkipCharFlag    int
	IgnoreFlag      bool
	InputType       string
	OutputType      string
}

func ArgsParser() (*Option, error){
	var newOption Option

	flag.BoolVar(&newOption.MutuallyExcFlag.CountFlag, "c", false, "count strings")
	flag.BoolVar(&newOption.MutuallyExcFlag.DuplicateFlag, "d", false, "only duplicate strings")
	flag.BoolVar(&newOption.MutuallyExcFlag.UniqueFlag, "u", false, "only unique strings")
	flag.IntVar(&newOption.SkipWordFlag, "f", 0, "skip first num_fildes words")
	flag.IntVar(&newOption.SkipCharFlag,"s", 0, "skip first num_fields chars")
	flag.BoolVar(&newOption.IgnoreFlag,"i", false, "ignore registr")

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println(" [input_file [output_file]]	filepath for input text and output text\n")
	}

	flag.Parse()


	args := flag.Args()
	switch len(args){
	case 0:
		newOption.InputType = os.Stdin.Name()
		newOption.OutputType = os.Stdout.Name()
	case 1:
		newOption.InputType = args[0]
		newOption.OutputType = os.Stdout.Name()
	case 2:
		newOption.InputType = args[0]
		newOption.OutputType = args[1]
	default:
		flag.Usage()
		return nil, errors.New("Wrong number of arguments.")
	}

	if newOption.MutuallyExcFlag.CountFlag {
		if newOption.MutuallyExcFlag.DuplicateFlag || newOption.MutuallyExcFlag.UniqueFlag {
			//error
			fmt.Fprintln(os.Stderr, "-c, -d, -u are mutually exclusive")
			flag.Usage()
			return nil, errors.New("Wrong arguments.")
		}
	} else {
		if newOption.MutuallyExcFlag.DuplicateFlag && newOption.MutuallyExcFlag.UniqueFlag {
			//error
			fmt.Fprintln(os.Stderr, "-c, -d, -u are mutually exclusive")
			flag.Usage()
			return nil, errors.New("Wrong arguments.")
		}
	}
	return &newOption, nil
}