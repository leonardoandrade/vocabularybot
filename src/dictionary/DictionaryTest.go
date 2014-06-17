package dictionary

import (
    "fmt"
)

func Test() {
    var d Dictionary = Make("test.json")
    d.Dump()
    var fixtures  = []string{"zeit", "fünf", "zukunft", "apfel"}

    for _,word := range fixtures {
        fmt.Println(word+" --> "+d.Lookup(word))
    }
}
