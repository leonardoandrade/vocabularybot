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
    spellCorrector SpellCorrector
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

    ret.spellCorrector= SpellCorrector{}

    return ret
}

func (d *Dictionary) Lookup(word string) (string) {
    return strings.Trim(d.dict[strings.ToLower(word)]," ")
}

func (d *Dictionary) Correct(word string) (string, string) {


    if meaning := d.dict[strings.ToLower(word)]; meaning != "" {
        return  word,meaning
    }
    corrections := d.spellCorrector.Correct(strings.ToLower(word))
    for _,w := range(corrections) {
        fmt.Println("w:"+w)
        if meaning := d.dict[strings.ToLower(w)]; meaning != "" {
            return  word,meaning
        }
    }


    return "", ""

}



func (d *Dictionary) Dump() {
    fmt.Println("dictionary size:", len(d.dict))
}
