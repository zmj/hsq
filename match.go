package main

type match struct {
	ourDecks     []*deck
	theirClasses []*class
}

func input(classes []*class) *match {
	return &match{
		ourDecks:     []*deck{classes[0].decks[0], classes[1].decks[1]},
		theirClasses: []*class{classes[2], classes[3]},
	}
}
