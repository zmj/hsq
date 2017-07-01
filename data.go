package main

func load() ([]*class, *matchups) {
	warrior := &class{name: "Warrior"}
	pirateWarrior := &deck{
		class: warrior,
		name:  "Pirate Warrior",
	}
	controlWarrior := &deck{
		class: warrior,
		name:  "Control Warrior",
	}
	warrior.decks = []*deck{pirateWarrior, controlWarrior}
	warrior.percentages = []int{80, 20}

	mage := &class{name: "Mage"}
	freezeMage := &deck{
		class: mage,
		name:  "Freeze Mage",
	}
	tempoMage := &deck{
		class: mage,
		name:  "Tempo Mage",
	}
	mage.decks = []*deck{freezeMage, tempoMage}
	mage.percentages = []int{50, 50}

	warlock := &class{name: "Warlock"}
	renoWarlock := &deck{
		class: warlock,
		name:  "Reno Warlock",
	}
	zooWarlock := &deck{
		class: warlock,
		name:  "Zoo Warlock",
	}
	warlock.decks = []*deck{renoWarlock, zooWarlock}
	warlock.percentages = []int{60, 40}

	druid := &class{name: "Druid"}
	rampDruid := &deck{
		class: druid,
		name:  "Ramp Druid",
	}
	eggDruid := &deck{
		class: druid,
		name:  "Egg Druid",
	}
	druid.decks = []*deck{rampDruid, eggDruid}
	druid.percentages = []int{30, 70}

	classes := []*class{warrior, mage, warlock, druid}
	matchups := &matchups{make(map[matchupKey]*matchup)}
	for _, c := range classes {
		for _, d := range c.decks {
			for _, c2 := range classes {
				for _, d2 := range c2.decks {
					matchups.add(d, d2, 50, 50)
				}
			}
		}
	}
	return classes, matchups
}
