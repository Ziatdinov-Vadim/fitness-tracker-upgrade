package trainings

import (
    "fmt"
    "strconv"
    "strings"
    "time"

    "github.com/Yandex-Practicum/tracker/internal/personaldata"
    "github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
    // TODO: добавить поля
    Steps        int
    TrainingType string
    Duration     time.Duration
    personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
    // TODO: реализовать функцию
    data := strings.Split(datastring, ",")
    if len(data) != 3 {
        return fmt.Errorf("invalid data format")
    }
    stepsStr := strings.TrimSpace(data[0])
    typeStr := strings.TrimSpace(data[1])
    durationStr := strings.TrimSpace(data[2])
    steps, err := strconv.Atoi(stepsStr)
    if err != nil || steps <= 0 {
        return fmt.Errorf("invalid steps format")
    }
    duration, err := time.ParseDuration(durationStr)
    if err != nil || duration <= 0 {
        return fmt.Errorf("invalid duration format")
    }
    t.Steps = steps
    t.TrainingType = typeStr
    t.Duration = duration
    return nil
}

func (t Training) ActionInfo() (string, error) {
    // TODO: реализовать функцию
    distance := spentenergy.Distance(t.Steps, t.Height)
    speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
    var calories float64
    var err error
    switch t.TrainingType {
    case "Бег":
        calories, err = spentenergy.RunningSpentCalories(
            t.Steps,
            t.Weight,
            t.Height,
            t.Duration,
        )
    case "Ходьба":
        calories, err = spentenergy.WalkingSpentCalories(
            t.Steps,
            t.Weight,
            t.Height,
            t.Duration,
        )
    default:
        return "", fmt.Errorf("неизвестный тип тренировки")
    }
    if err != nil {
        return "", err
    }
    result := fmt.Sprintf(
        "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
        t.TrainingType,
        t.Duration.Hours(),
        distance,
        speed,
        calories,
    )
    return result, nil
}