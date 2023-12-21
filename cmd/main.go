package main

import (
	"flag"
	"fmt"

	"github.com/nearsyh/regex/regex"
)

var (
	str string
	pat string
	debug bool
)

func main() {
	str := flag.String("str", "", "The string to be matched")
	pat := flag.String("pat", "", "The regex pattern")
	debug := flag.Bool("debug", false, "If it is set as true, a dot file of the regex state machine would be generated")
	flag.Parse()
	if *debug {
		regex.Visualize(*pat, "./debug.dot")
	}
	fmt.Println(regex.Match(*str, *pat))
}
