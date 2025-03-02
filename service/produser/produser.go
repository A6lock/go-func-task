package produser

import (
	"bufio"
	"fmt"
	"os"
)

type Producer struct {
	inputFilePath string
}

func NewProducer(inputFilePath string) *Producer {
	return &Producer{inputFilePath}
}

func (p Producer) Produce() ([]string, error) {
	file, err := os.Open(p.inputFilePath)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла %s", p.inputFilePath)

		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при чтении файла %s", p.inputFilePath)
		return nil, err
	}

	if len(lines) == 0 {
		fmt.Printf("Файл %s пуст\n", p.inputFilePath)
		return nil, fmt.Errorf("файл %s пуст", p.inputFilePath)
	}

	return lines, nil
}
