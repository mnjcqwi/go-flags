package main

import (
	"flag"
	"fmt"
)

func main() {
	color := flag.String("fg", "000000", "background color")
	flag.Parse()

	fmt.Println(*color)
}
