package main

import (
	"fmt"
	"os"
)

//docker             run image <cmd> <params>
// go run main.go    run       <cmd> <params>

func main() {
	switch os.Args[1] {
	case "run":
		run()

	default:
		panic("bad command")
	}
}

func run() {
	fmt.Print("Running ", os.Args[2:])
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
