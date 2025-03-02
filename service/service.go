package service

import (
	"fmt"
	"sync"
)

const prefix = "http://"
const prefixLength = len(prefix)

type Producer interface {
	Produce() ([]string, error)
}

type Presenter interface {
	Present([]string) error
}

type Service struct {
	prod Producer
	pres Presenter
}

func NewService(prod Producer, pres Presenter) *Service {
	return &Service{prod, pres}
}

func (s Service) maskingLinks(inputStr string) string {
	isLink := false

	strByteSlice := []byte(inputStr)

	for i := 0; i < len(strByteSlice); i++ {
		if string(strByteSlice[i]) == " " {
			isLink = false
		}

		if isLink {
			strByteSlice[i] = '*'

			continue
		}

		if i < len(strByteSlice)-prefixLength && string(inputStr[i:i+prefixLength]) == prefix {
			isLink = true

			i += prefixLength - 1
		}
	}

	return string(strByteSlice)
}

func (s Service) Run() error {
	semaphore := make(chan struct{}, 10)
	var wg sync.WaitGroup

	inputSlice, err := s.prod.Produce()
	if err != nil {
		return fmt.Errorf("ошибка при чтении файла: %v", err)
	}

	resultChan := make(chan struct {
		index int
		text  string
	}, len(inputSlice))

	for i, line := range inputSlice {
		wg.Add(1)
		go func(index int, text string) {
			defer wg.Done()

			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			maskedText := s.maskingLinks(text)

			resultChan <- struct {
				index int
				text  string
			}{index, maskedText}
		}(i, line)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	maskedSlice := make([]string, len(inputSlice))

	for result := range resultChan {
		maskedSlice[result.index] = result.text
	}

	err = s.pres.Present(maskedSlice)
	if err != nil {
		return fmt.Errorf("ошибка при записи файла: %v", err)
	}

	return nil
}
