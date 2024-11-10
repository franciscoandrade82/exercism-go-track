package dndcharacter

import (
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	modifier := score - 10
	if modifier%2 != 0 && modifier < 0 {
		modifier -= 1
	}
	return modifier / 2
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	min := 6
	sum := 0
	for i := 0; i < 4; i++ {
		roll := rand.Intn(6) + 1
		sum += roll
		if roll <= min {
			min = roll
		}
	}

	return sum - min
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	character := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}

	character.Hitpoints = 10 + Modifier(character.Constitution)

	return character
}
