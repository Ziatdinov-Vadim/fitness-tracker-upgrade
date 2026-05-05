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
    // TODO: реализовать функцию
    if steps <= 0 {
        return 0, fmt.Errorf("invalid steps")
    }
    if weight <= 0 {
        return 0, fmt.Errorf("invalid weight")
    }
    if height <= 0 {
        return 0, fmt.Errorf("invalid height")
    }
    if duration <= 0 {
        return 0, fmt.Errorf("invalid duration")
    }
    speed := MeanSpeed(steps, height, duration)
    calories := weight *
        walkingCaloriesCoefficient *
        speed *
        float64(duration.Hours())
    return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
    // TODO: реализовать функцию
    if steps <= 0 {
        return 0, fmt.Errorf("invalid steps")
    }
    if weight <= 0 {
        return 0, fmt.Errorf("invalid weight")
    }
    if height <= 0 {
        return 0, fmt.Errorf("invalid height")
    }
    if duration <= 0 {
        return 0, fmt.Errorf("invalid duration")
    }
    speed := MeanSpeed(steps, height, duration)
    const runningCaloriesCoefficient = 1.0
    calories := weight * speed * runningCaloriesCoefficient * duration.Hours()
    return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
    // TODO: реализовать функцию
    if steps <= 0 || height <= 0 || duration <= 0 {
        return 0
    }
    distance := Distance(steps, height)
    return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
    // TODO: реализовать функцию
    if steps <= 0 || height <= 0 {
        return 0
    }
    stepLength := height * stepLengthCoefficient
    distanceMeters := float64(steps) * stepLength
    return distanceMeters / mInKm
}
