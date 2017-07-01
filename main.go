package main

import (
	"fmt"
)

func main() {
	classes, matchups := load()
	match := input(classes)

	fmt.Printf("%v\n", matchups)
	worlds := match.worlds()
	for _, w := range worlds {
		fmt.Printf("probability: %v\n", w.probability)
		fmt.Println("ours")
		for _, d := range w.ourDecks {
			fmt.Printf("%v\n", d)
		}
		fmt.Println("theirs")
		for _, d := range w.theirDecks {
			fmt.Printf("%v\n", d)
		}
		fmt.Println()
	}
}

type world struct {
	probability int
	ourDecks    []*deck
	theirDecks  []*deck
}

func (m *match) worlds() []*world {
	var rec func(classes []*class) []*world
	rec = func(classes []*class) []*world {
		if len(classes) == 0 {
			return []*world{&world{probability: 1}}
		}
		c := classes[0]
		subworlds := rec(classes[1:])
		var worlds []*world
		for i, d := range c.decks {
			for _, sub := range subworlds {
				w := &world{
					probability: sub.probability * c.percentages[i],
					ourDecks:    m.ourDecks,
					theirDecks:  append(sub.theirDecks, d),
				}
				worlds = append(worlds, w)
			}
		}
		return worlds
	}
	return rec(m.theirClasses)
}
