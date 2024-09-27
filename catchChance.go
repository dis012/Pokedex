package main

import (
	"math/rand"
	"time"
)

const maxBaseExperience = 635

func calculateCatchChance(base_experience int) float64 {
	difficultyFactor := float64(base_experience / maxBaseExperience)

	luckFactor := 0.7 + rand.Float64()*0.3

	catchChance := (1.0 - difficultyFactor) * luckFactor

	if catchChance < 0 {
		return 0.0
	}
	return catchChance
}

func attemptCatch(base_experience int) bool {
	catchChance := calculateCatchChance(base_experience)

	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Float64()

	return randomNumber <= catchChance
}
