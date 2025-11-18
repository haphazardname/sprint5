package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// расчет потраченных калорий при хотьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные входные параметры WalkingSpentCalories")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return ((weight * meanSpeed * durationInMinutes) / float64(minInH)) * float64(walkingCaloriesCoefficient), nil
}

// расчет потраченных калорий при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные входные параметры RunningSpentCalories")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / float64(minInH), nil
}

// расссчет средней скорости
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration > 0 && steps > 0 && height > 0 {
		return Distance(steps, height) / duration.Hours()
	}
	return 0
}

// рассчет дистанции в километрах
func Distance(steps int, height float64) float64 {
	if steps > 0 && height > 0 {
		return ((float64(stepLengthCoefficient) * height) * float64(steps)) / float64(mInKm)
	}
	return 0
}
