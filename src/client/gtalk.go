package client

import (
        "fmt"
        "os"
        "log"
      	"client/xmpp"
)


const (
	VOCABULARYBOT_USERNAME = "VOCABULARYBOT_USERNAME"
	VOCABULARYBOT_PASSWORD = "VOCABULARYBOT_PASSWORD"
	SERVER                 = "talk.google.com:443"
	NOTLS                  = false
	DEBUG                  = true
	SESSION                = true
)

type Message struct {

}



func (c *Client) InitGtalk() {
	var username = os.Getenv(VOCABULARYBOT_USERNAME)
	if username == "" {
		fmt.Printf("variable '%v' must be set\n", VOCABULARYBOT_USERNAME)
		os.Exit(0)
	}

	var password = os.Getenv(VOCABULARYBOT_PASSWORD)
	if password == "" {
		fmt.Printf("variable '%v' must be set\n", VOCABULARYBOT_PASSWORD)
		os.Exit(0)
	}

	var talk *xmpp.Client
	var err error
	options := xmpp.Options{Host: SERVER,
		User:     username,
		Password: password,
		NoTLS:    NOTLS,
		Debug:    DEBUG,
		Session:  SESSION}

  fmt.Println("options %v", options)
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
                c.Output <- v.Text
                response := <- c.Input
				fmt.Println("received text: " + v.Text + " dict lookup:" + response)

				if v.Text != "" {
					var msg = xmpp.Chat{
						Remote: v.Remote,
						Type:   "chat",
						Text:   response}
					talk.Send(msg)
				}
			case xmpp.Presence:
				fmt.Println(v.From, v.Show)
			}
		}
	}()
}
