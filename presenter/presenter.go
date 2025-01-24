package presenter

import (
	"fmt"
	"os"
	"strings"
)

const defaultPath = "output.txt"

type Presenter struct {
	outputFilePath string
}

func NewPresenter(outputFilePath string) Presenter {
	return Presenter{
		outputFilePath,
	}
}

func (p Presenter) Present(inputSlice []string) error {
	file, err := os.Create(p.outputFilePath)
	if err != nil {
		fmt.Printf("Ошибка при создании файла по пути: %s", p.outputFilePath)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(strings.Join(inputSlice, "\n"))
	if err != nil {
		fmt.Println("Ошибка при записи файла")
		return err
	}

	return nil
}
