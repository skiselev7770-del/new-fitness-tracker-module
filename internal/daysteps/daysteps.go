package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
	if len(parts) < 2 {
		return fmt.Errorf("ожидалось элементов: 2, получено: %d", len(parts))
	}
	if len(parts) > 2 {
		return fmt.Errorf("ожидалось элементов: 2, получено: %d", len(parts))
	}

	sumSteps, err := strconv.Atoi(parts[0])
	if sumSteps <= 0 {
		return fmt.Errorf("количество шагов не может быть меньше нуля, получено: %d", sumSteps)
	}
	if err != nil {
		return fmt.Errorf("не удалось преобразовать %s в число шагов: %w", parts[0], err)
	}

	ds.Steps = sumSteps

	durationWalk, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("не удалось преобразовать %s в длительность: %w", parts[1], err)
	}
	if durationWalk <= 0 {
		return fmt.Errorf("продолжительность прогулки не может быть меньше нуля, получено: %v", durationWalk)
	}

	ds.Duration = durationWalk

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("Ошибка: %w", err)
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		spentCalories), nil
}
