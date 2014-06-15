/*
* Based on example.go from github.com/mattn/go-xmpp
 */

package main

import (
	"bufio"
	"fmt"
	"github.com/mattn/go-xmpp"
	"log"
	"os"
	"strings"
)

const (

    VOCABULARYBOT_USERNAME = "VOCABULARYBOT_USERNAME"
    VOCABULARYBOT_PASSWORD = "VOCABULARYBOT_PASSWORD"
)
var server = "talk.google.com:443"
var notls = false
var debug = true
var session = false

func main() {

    var username = os.Getenv(VOCABULARYBOT_USERNAME)
    if username == "" {
        fmt.Printf("variable '%v' must be set\n", VOCABULARYBOT_USERNAME)
        return
    }

    var password = os.Getenv(VOCABULARYBOT_PASSWORD)
    if password == "" {
        fmt.Printf("variable '%v' must be set\n", VOCABULARYBOT_PASSWORD)
        return
    }


	var talk *xmpp.Client
	var err error
	options := xmpp.Options{Host: server,
		User:     username,
		Password: password,
		NoTLS:    notls,
		Debug:    debug,
		Session:  session}

	talk, err = options.NewClient()

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			chat, err := talk.Recv()
			if err != nil {
				log.Fatal(err)
			}
			switch v := chat.(type) {
			case xmpp.Chat:
				fmt.Println(v.Remote, v.Text, v.Type, v.Other)
				if v.Text != "" {
					var msg = xmpp.Chat{
						Remote: v.Remote,
						Type:   "chat",
						Text:  strings.ToUpper(v.Text)}
					talk.Send(msg)
				}
			case xmpp.Presence:
				fmt.Println(v.From, v.Show)
			}
		}
	}()
	for {
		in := bufio.NewReader(os.Stdin)
		line, err := in.ReadString('\n')
		if err != nil {
			continue
		}
		line = strings.TrimRight(line, "\n")

		tokens := strings.SplitN(line, " ", 2)
		if len(tokens) == 2 {
			talk.Send(xmpp.Chat{Remote: tokens[0], Type: "chat", Text: tokens[1]})
		}
	}
}