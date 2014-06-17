package dictionary

import (
    "fmt"
    "strings"
    "os"
    "bufio"
    "encoding/json"
)

type  Dictionary struct {
    dict map[string]string
}

func  Make(fileName string) (Dictionary) {
    //TODO
    var ret Dictionary = Dictionary{dict: make(map[string]string)}

    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println(err)
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        var line = scanner.Text()
        var e Entry = Entry{}
        json.Unmarshal([]byte(line), &e)
        ret.dict[strings.ToLower(e.Key)] = strings.Join(e.Meanings, ", ")
    }

    return ret

}

func (d *Dictionary) Lookup(word string) (string) {
    //TODO
    return d.dict[strings.ToLower(word)]
}

func (d *Dictionary) Dump() {
    fmt.Println("dictionary size:", len(d.dict))
}
