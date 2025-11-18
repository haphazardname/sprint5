package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {

	if len(dataset) == 0 {
		return
	}

	for _, datastring := range dataset {
		err := dp.Parse(datastring)
		if err != nil {
			log.Println("Ошибка парсинга данных:", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("Ошибка при получении информации:", err)
			continue
		}
		fmt.Println(info)
	}

}
