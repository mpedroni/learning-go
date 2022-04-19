package main

import "fmt"

type Speaker interface {
	SayEffectPhrase()
}

type Character struct {
	AlterEgo, Name, City, EffectPhrase string
}

func (c Character) SayEffectPhrase() {
	fmt.Printf("%v: %v\n", c.Name, c.EffectPhrase)
}

type SuperHero struct {
	Character
	SuperPower string
}

type Villain struct {
	Character
	ArchEnemy SuperHero
}

func EnterTheScene(s Speaker) {
	s.SayEffectPhrase()
}

func main() {
	batman := SuperHero{
		Character: Character{
			AlterEgo:     "Bruce Wayne",
			Name:         "Batman",
			City:         "Gotham",
			EffectPhrase: "I'm Batman",
		},
		SuperPower: "Money",
	}

	joker := Villain{
		Character: Character{
			AlterEgo:     "Jack Napier",
			Name:         "Joker",
			City:         "Gotham",
			EffectPhrase: "Why so serious?",
		},
		ArchEnemy: batman,
	}

	EnterTheScene(batman)
	// or
	batman.SayEffectPhrase()

	EnterTheScene(joker)
	// or
	joker.SayEffectPhrase()

	// batman and joker can call SayEffectPhrase because they are both Character too

}
