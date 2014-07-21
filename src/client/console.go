package client


import (
	"fmt"
)

func (c *Client) InitConsole() {

	var word string

	for {

		fmt.Scanf("%s", &word)
		c.Output <- word
		response := <- c.Input
		fmt.Println("--> ",response)
	}

}
