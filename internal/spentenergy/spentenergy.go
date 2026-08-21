package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов не может быть меньше нуля, получено: %d", steps)
	}
	if weight <= 0 || height <= 0 {
		return 0, fmt.Errorf("вес и рост не могут быть меньше нуля, получено: %.2f и %.2f", weight, height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность прогулки не может быть меньше нуля, получено: %v", duration)
	}

	averageSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := (weight * averageSpeed * durationInMinutes) / minInH

	return spentCalories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов не может быть меньше нуля, получено: %d", steps)
	}
	if weight <= 0 || height <= 0 {
		return 0, fmt.Errorf("вес и рост не могут быть меньше нуля, получено: %.2f и %.2f", weight, height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность прогулки не может быть меньше нуля, получено: %v", duration)
	}

	averageSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := (weight * averageSpeed * durationInMinutes) / minInH

	return spentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	averageSpeed := distance / duration.Hours()

	return averageSpeed
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	stepLength := height * stepLengthCoefficient
	sumSteps := float64(steps) * stepLength
	distance := sumSteps / mInKm

	return distance
}
