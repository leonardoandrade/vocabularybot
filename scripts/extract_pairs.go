package main

import (
	"fmt"
	"regexp"
    "os"
    "bufio"
)

/* TODO:
* - threat plurals
* - extract gender
* - other corner cases to treat
*/


var regexLine  = regexp.MustCompile("^([A-Za-zöäü ]+) \\{([a-z]+)\\}.+::([A-Za-z;\\-\\(\\) ]+)|.*$")

func main() {
    file, _ := os.Open("../data/de-en.txt")

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        var line = scanner.Text()

        var matches = regexLine.FindStringSubmatch(line)

	    if len(matches) == 4 && matches[1] != "" {
		    fmt.Printf("%s;%s;%s\n", matches[1], matches[2], matches[3])
	    }
    }
}
