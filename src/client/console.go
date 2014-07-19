package client


import (
	"fmt"
)

func InitConsole(c *Client) {

	var word string

	for {

		fmt.Scanf("%s", &word)

		c.Output <- word
		response := <- c.Input
		fmt.Println("--> ",response)
	}

}
