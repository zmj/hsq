package main

import (
	"fmt"
)

func main() {
	classes, matchups := load()
	for _, c := range classes {
		for _, d := range c.decks {
			fmt.Printf("%v\n", d)
		}
	}
	for _, m := range matchups.m {
		fmt.Printf("%v\n", m)
	}
}
