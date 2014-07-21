package main

import (
	"client"
	"dictionary"
	"fmt"
	"os"
	"time"
)

func usage() {
	fmt.Println(" usage:", os.Args[0], " <client-type> <json dict file>")
	fmt.Println(" where <client-type> in [console, gtalk]")
}

func main() {

	if len(os.Args) != 3 {
		usage()
		return
	}

	var dict dictionary.Dictionary = dictionary.Make(os.Args[2])

	cl := client.Make()

	switch os.Args[1] {
	case "console":
		go cl.InitConsole()
	case "gtalk":
		go cl.InitGtalk()
	default:
		usage()
		return
	}

	count := 1
	for {

		req := <-cl.Output
		if count == 1 {
			time.Sleep(time.Duration(5) * time.Second)
		}
		response := dict.Lookup(req)
		if response == "" {
			correction, meaning := dict.Correct(req)
			if correction != "" {
				response = "did you mean '" + correction + "'? that means '" + meaning + "'"
			} else {
				response = "word not found :-("
			}
		}
		fmt.Println(response)
		cl.Input <- response

		count = count + 1
	}
}
