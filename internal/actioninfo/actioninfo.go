package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for i, line := range dataset {
		err := dp.Parse(line)
		if err != nil {
			log.Printf("ошибка парсинга на строке %d: %v", i, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка формирования информации об активности: %v\n", err)
			return
		}
		fmt.Println(info)
	}

}
