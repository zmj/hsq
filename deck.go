package main

import (
	"fmt"
)

type deck struct {
	class *class
	name  string
}

type matchup struct {
	deck1, deck2       *deck
	deck1win, deck2win int
}

type class struct {
	name        string
	decks       []*deck
	percentages []int
}

type matchupKey struct {
	deck1, deck2 *deck
}

type matchups struct {
	m map[matchupKey]*matchup
}

func (m *matchups) add(d1, d2 *deck, w1, w2 int) {
	matchup := &matchup{d1, d2, w1, w2}
	m.m[matchupKey{d1, d2}] = matchup
	m.m[matchupKey{d2, d1}] = matchup
}

func (m *matchup) String() string {
	return fmt.Sprintf("%v (%v) vs %v (%v)", m.deck1, m.deck1win, m.deck2, m.deck2win)
}

func (d *deck) String() string {
	return d.name
}

func (c *class) String() string {
	return c.name
}
