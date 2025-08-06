package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println()

	name := flag.String("name", "default", "a name to greet")

	flag.Parse()

	fmt.Println(*name)
}
