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
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
	if len(parts) < 3 {
		return fmt.Errorf("ожидалось элементов: 3, получено: %d", len(parts))
	}
	if len(parts) > 3 {
		return fmt.Errorf("ожидалось элементов: 3, получено: %d", len(parts))
	}

	sumSteps, err := strconv.Atoi(parts[0])
	if sumSteps <= 0 {
		return fmt.Errorf("количество шагов не может быть меньше нуля, получено: %d", sumSteps)
	}
	if err != nil {
		return fmt.Errorf("не удалось преобразовать %s в число шагов: %w", parts[0], err)
	}

	t.Steps = sumSteps

	t.TrainingType = parts[1]

	durationWalk, err := time.ParseDuration(parts[2])

	if err != nil {
		return fmt.Errorf("не удалось преобразовать %s в длительность: %w", parts[2], err)
	}
	if durationWalk <= 0 {
		return fmt.Errorf("продолжительность прогулки не может быть меньше нуля, получено: %v", durationWalk)
	}

	t.Duration = durationWalk
	return nil

}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Height)
	averageSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var spentCalories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("ошибка расчёта калорий для типа %q: %w", t.TrainingType, err)
		}
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("ошибка расчёта калорий для типа %q: %w", t.TrainingType, err)
		}
	default:

		return "", fmt.Errorf("неподдерживаемый тип тренировки: %q", t.TrainingType)
	}

	return fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType,
			t.Duration.Hours(),
			distance,
			averageSpeed,
			spentCalories),
		nil
}
