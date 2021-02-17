package read_write

import "testing"

func TestReadText(t *testing.T) {
	_, err := ReadText("dsdsd")
	if err == nil{
		t.Error("Error was expected.")
	}
}

func TestReadText2(t *testing.T) {
	_, err := ReadText("../input.txt")
	if err != nil{
		t.Error("Unexpected error.")
	}
}

func TestWriteText(t *testing.T) {
	testStrings := []string{"sdjskhj", "dsjbjhk"}
	err := WriteText(testStrings, "/weuyh")
	if err == nil{
		t.Error("Error was expected.")
	}
}

func TestWriteText2(t *testing.T) {
	testStrings := []string{"sdjskhj", "dsjbjhk"}
	err := WriteText(testStrings, "../output.txt")
	if err != nil{
		t.Error("Unexpected error.")
	}
}
