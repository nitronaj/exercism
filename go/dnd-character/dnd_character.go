package dndcharacter

import (
	"math"
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
	return int(math.Floor(float64(score-10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	dices := make([]int, 0, 4)
	sum, minIndex := 0, 0

	for i := 0; i < 4; i++ {
		roll := rand.Intn(6) + 1
		dices = append(dices, roll)
		if dices[i] < dices[minIndex] {
			minIndex = i
		}
		sum += roll
	}

	sum -= dices[minIndex]

	return sum
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	character := Character{}
	character.Strength = Ability()
	character.Dexterity = Ability()
	character.Constitution = Ability()
	character.Intelligence = Ability()
	character.Wisdom = Ability()
	character.Charisma = Ability()
	character.Hitpoints = 10 + Modifier(character.Constitution)

	return character
}
