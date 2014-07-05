package client

type Client struct {
    Input  chan string
    Output chan string
}


func Make() (Client){
    ret := Client{Input : make(chan string), Output : make(chan string)}
    return ret
}
