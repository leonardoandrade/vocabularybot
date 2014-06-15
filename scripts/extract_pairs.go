package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type DictionaryEntry struct {
	Key      string
	TermType string
	Meanings []string
}

var regexLine = regexp.MustCompile("^([A-Za-zöäü ]+) \\{([a-z]+)\\}.+::([A-Za-z;\\-\\(\\) ]+)|.*$")

func writeJSONToFile(dict []DictionaryEntry, fileName string) (bool, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return false, err
	}

	defer func() {
		file.Close()
	}()

	w := bufio.NewWriter(file)
	for _, entry := range dict {
		b, err := json.Marshal(entry)
		if err != nil {
			return false, err
		}

		w.WriteString(string(b) + "\n")
	}
	return true, nil
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("usage:", os.Args[0], " <src raw dictionary file> <dst json file>")
		return
	}
	var srcFile string = os.Args[1]
	var dstFile string = os.Args[2]

	file, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	dict := []DictionaryEntry{}
	termSet := make(map[string]bool)

	for scanner.Scan() {
		var line = scanner.Text()

		var matches = regexLine.FindStringSubmatch(line)

		key := strings.Trim(matches[1], " ")

		//do not include repeated terms
		if termSet[key] {
			continue
		}

		termSet[key] = true

		var meanings []string = strings.Split(matches[3], ";")

		if strings.Trim(matches[1], " ") != "" {
			entry := DictionaryEntry{Key: key, TermType: matches[2], Meanings: meanings}
			dict = append(dict, entry)
		}
	}

	_, err1 := writeJSONToFile(dict, dstFile)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	fmt.Println(len(dict), "entries written to file", dstFile)
}
