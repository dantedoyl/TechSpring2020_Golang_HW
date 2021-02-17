package uniq

import (
	"fmt"
	"strings"
	ap "unique/args_parser"
)

type uniqString struct {
	str      string
	origStr  string
	quantity int
}

func searchString(uniqStrings []uniqString, str string) int {
	for ind, el := range uniqStrings {
		if el.str == str {
			return ind
		}
	}
	return -1
}

func transformStrings(text []string, skipWord int, skipChar int, ignoreCase bool) []uniqString {
	var uniqStrings []uniqString

	for _, str := range text {
		var newStr string

		switch {
		case len(strings.Split(str, " ")) < skipWord:
			newStr = ""
		case len(strings.Split(str, " ")) == skipWord:
			newStr = str
		default:
			newStr = strings.Join(strings.Split(str, " ")[skipWord:], " ")
		}

		if len(newStr) < skipChar {
			newStr = ""
		} else {
			newStr = newStr[skipChar:]
		}

		if ignoreCase {
			newStr = strings.ToLower(newStr)
		}

		ind := searchString(uniqStrings, newStr)
		if ind == -1{
			uniqStrings = append(uniqStrings,uniqString{
				str:      newStr,
				origStr:  str,
				quantity: 1,
			})
		} else {
			uniqStrings[ind].quantity++
		}
	}
	return uniqStrings
}

func Uniq(text []string, option ap.Option) ([]string){
	resultSlice  := []string{}

	strSlice := transformStrings(text, option.SkipWordFlag, option.SkipCharFlag, option.IgnoreFlag)

	switch {
	case option.MutuallyExcFlag.CountFlag:
		for _, str := range strSlice{
			resultSlice = append(resultSlice, fmt.Sprintf("%d %s", str.quantity, str.origStr))
		}
	case option.MutuallyExcFlag.DuplicateFlag:
		for _, str := range strSlice{
			if(str.quantity > 1) {
				resultSlice = append(resultSlice, str.origStr)
			}
		}
	case option.MutuallyExcFlag.UniqueFlag:
		for _, str := range strSlice{
			if(str.quantity == 1) {
				resultSlice = append(resultSlice, str.origStr)
			}
		}
	default:
		for _, str := range strSlice{
			resultSlice = append(resultSlice, str.origStr)
		}
	}

	return resultSlice
}
