package uniq

import (
	"reflect"
	"testing"
	ap "unique/args_parser"
)

func TestUniq(t *testing.T) {
	testStrings := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	expectedStrings := []string{
		"3 I love music.",
		"1 ",
		"2 I love music of Kartik.",
		"1 Thanks.",
	}

	mexFlag := ap.MutuallyExclusiveFlag{
		CountFlag:     true,
		DuplicateFlag: false,
		UniqueFlag:    false,
	}

	testOption := ap.Option{
		MutuallyExcFlag: mexFlag,
		SkipWordFlag:    0,
		SkipCharFlag:    0,
		IgnoreFlag:      true,
		InputType:       "",
		OutputType:      "",
	}

	testResult := Uniq(testStrings, testOption)
	if !reflect.DeepEqual(testResult, expectedStrings){
		t.Error("Wrong answer.")
	}

}

func TestUniq1(t *testing.T) {
	testStrings := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	expectedStrings := []string{
		"I love music.",
	}

	mexFlag := ap.MutuallyExclusiveFlag{
		CountFlag:     false,
		DuplicateFlag: false,
		UniqueFlag:    false,
	}

	testOption := ap.Option{
		MutuallyExcFlag: mexFlag,
		SkipWordFlag:    10,
		SkipCharFlag:    20,
		IgnoreFlag:      false,
		InputType:       "",
		OutputType:      "",
	}

	testResult := Uniq(testStrings, testOption)
	if !reflect.DeepEqual(testResult, expectedStrings){
		t.Error("Wrong answer.")
	}

}

func TestUniq2(t *testing.T) {
	testStrings := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	expectedStrings := []string{
		"I love music.",
		"I love music of Kartik.",
	}

	mexFlag := ap.MutuallyExclusiveFlag{
		CountFlag:     false,
		DuplicateFlag: true,
		UniqueFlag:    false,
	}

	testOption := ap.Option{
		MutuallyExcFlag: mexFlag,
		SkipWordFlag:    1,
		SkipCharFlag:    0,
		IgnoreFlag:      false,
		InputType:       "",
		OutputType:      "",
	}

	testResult := Uniq(testStrings, testOption)
	if !reflect.DeepEqual(testResult, expectedStrings){
		t.Error("Wrong answer.")
	}

}

func TestUniq3(t *testing.T) {
	testStrings := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	expectedStrings := []string{
		"",
		"Thanks.",
	}

	mexFlag := ap.MutuallyExclusiveFlag{
		CountFlag:     false,
		DuplicateFlag: false,
		UniqueFlag:    true,
	}

	testOption := ap.Option{
		MutuallyExcFlag: mexFlag,
		SkipWordFlag:    0,
		SkipCharFlag:    0,
		IgnoreFlag:      false,
		InputType:       "",
		OutputType:      "",
	}

	testResult := Uniq(testStrings, testOption)
	if !reflect.DeepEqual(testResult, expectedStrings){
		t.Error("Wrong answer.")
	}

}
