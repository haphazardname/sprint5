package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

type DaySteps struct {
	Steps                 int           // количество шагов.
	Duration              time.Duration // длительность прогулки.
	personaldata.Personal               // встроенная структура Personal из пакета personaldata, у которой есть метод Print().
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dataSlice := strings.Split(datastring, ",")
	if len(dataSlice) != 2 {
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
	ds.Steps = steps

	trainDur, err := time.ParseDuration(dataSlice[1])
	if err != nil || trainDur.Seconds() <= 0 {
		err := errors.New("ошибка преобразования длительности тренировки parsePackage")
		log.Println(err)
		return err
	}
	ds.Duration = trainDur
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {

	if ds.Steps <= 0 || ds.Duration.Seconds() <= 0 {
		err := errors.New("неверный формат входных данных ActionInfo")
		log.Println(err)
		return "", err
	}
	distance := (stepLength * float64(ds.Steps)) / float64(mInKm)
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, trainDur)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)

}
