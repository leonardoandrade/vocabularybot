/*
* Based on example.go from github.com/mattn/go-xmpp
 */

package main

import (
	"fmt"
	"os"
    "./dictionary"
    "./client/"
)



func main() {

    if len(os.Args) != 2 {
        fmt.Println("usage:", os.Args[0], " <json dict file>")
        return
    }

    var dict dictionary.Dictionary  = dictionary.Make(os.Args[1])
    gtalk := client.Make()

    go gtalk.Init()

	for {
	       req := <- gtalk.Output
           response := dict.Lookup(req)
           if response == "" {
               response = "word not found :-("
           }
           fmt.Println(response)
           gtalk.Input <- response
	}
}
