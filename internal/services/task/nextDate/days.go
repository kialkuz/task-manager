package nextDate

import (
	"errors"
	"strconv"

	"github.com/kialkuz/task-manager/internal/infrastructure/env"
)

func checkIntervalTypeDays(formatParts []string) error {
	if len(formatParts) < 2 || formatParts[1] == "" {
		return errors.New("неверный формат интервала")
	}

	count, err := strconv.Atoi(formatParts[1])
	if err != nil {
		return errors.New("количество должно быть цифрой")
	}

	if count <= 0 {
		return errors.New("передан нулевой или отрицательный интервал")
	}

	maxIntervalDaysNumber, err := strconv.Atoi(env.GetEnv("TODO_MAX_INTERNAL_DAYS", ""))
	if err != nil {
		return errors.New("максимально допустимое число дней должно быть цифрой")
	}

	if count > maxIntervalDaysNumber {
		return errors.New("превышен максимально допустимый интервал")
	}

	return nil
}
