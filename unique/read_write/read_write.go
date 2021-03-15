package read_write

import (
	"bufio"
	"os"
)

func ReadText(fileName string) ([]string, error) {
	var text []string

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileString := bufio.NewScanner(file)

	for fileString.Scan() {
		text = append(text, fileString.Text())
	}

	return text, nil
}

func WriteText(text []string, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer f.Close()

	for _, str := range text {
		_, err := f.WriteString(str + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
