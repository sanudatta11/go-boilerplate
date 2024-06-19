package main

import (
	"boilerplate/server"
	"flag"
	"fmt"
	"os"
)

// @schemes http
func main() {
	env := flag.String("e", "", "Specify the environment (default is development)")
	flag.Usage = func() {
		fmt.Printf("Usage: %s -e <environment>\n", os.Args[0])
		os.Exit(1)
	}

	flag.Parse()
	fmt.Printf("Using Env %v\n", *env)
	server.Init(*env)
}
