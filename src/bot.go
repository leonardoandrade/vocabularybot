package main

import (
	"fmt"
	"os"
    "./dictionary"
    "./client"
	"./client/gtalk"
    "time"
)



func main() {

    if len(os.Args) != 2 {
        fmt.Println("usage:", os.Args[0], " <json dict file>")
        return
    }

    var dict dictionary.Dictionary  = dictionary.Make(os.Args[1])
    cl := client.Make()

    go gtalk.Init(&cl)

    count :=1
	for {

	       req := <- cl.Output
           if count ==1 {
               time.Sleep(time.Duration(5) * time.Second)
           }
           response := dict.Lookup(req)
           if response == "" {
               correction, meaning := dict.Correct(req)
               fmt.Println("CORRECTIONL"+correction)
               if correction != "" {
                   response = "did you mean '"+correction+"'? that means '"+meaning+"'"
               } else {
                   response = "word not found :-("
               }
           }
           fmt.Println(response)
           cl.Input <- response

           count = count + 1
	}
}
