package main

import "fmt"

type ServerState int

const (
	Idle = iota
	Connected
	Error
	Retrying
)

var stateName = map[ServerState]string{
	Idle:      "idle",
	Connected: "connected",
	Error:     "error",
	Retrying:  "retrying",
}

func (ss ServerState) toString() string {
	return stateName[ss]
}

func main() {
	var state ServerState = Idle
	fmt.Println(state.toString())
}
