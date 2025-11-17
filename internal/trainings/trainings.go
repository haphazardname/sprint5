package trainings

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps                 int           //количество шагов, проделанных за тренировку.
	TrainingType          string        //тип тренировки(бег или ходьба).
	Duration              time.Duration //длительность тренировки.
	personaldata.Personal               // встроенная структура Personal из пакета personaldata, у которой есть метод Print().
}

func (t *Training) Parse(datastring string) (err error) {
	dataSlice := strings.Split(datastring, ",")
	if len(dataSlice) != 3 {
		err := errors.New("неверный формат входных данных parsePackage")
		log.Println(err)
		return err
	}

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil || steps <= 0 {
		err := errors.New("ошибка преобразования шагов parsePackage")
		log.Println(err)
		return err
	}
	t.Steps = steps
	t.TrainingType = dataSlice[1]

	trainDur, err := time.ParseDuration(dataSlice[2])
	if err != nil || trainDur.Seconds() <= 0 {
		err := errors.New("ошибка преобразования длительности тренировки parsePackage")
		log.Println(err)
		return err
	}
	t.Duration = trainDur
	return nil
}

func (t Training) ActionInfo() (string, error) {

	if t.Steps <= 0 || t.Height <= 0 || t.Duration.Seconds() <= 0 {
		return "", errors.New("некорректные данные для тренировки")
	}

	var calories float64
	var err error
	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		err = errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2fч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps, t.Personal.Height), spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration), calories)

	return info, nil

}
