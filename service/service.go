package service

import (
	"fmt"
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

func NewService(prod Producer, pres Presenter) Service {
	return Service{prod, pres}
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
	inputSlice, err := s.prod.Produce()
	if err != nil {
		return fmt.Errorf("ошибка при чтении файла: %v", err)
	}

	maskedSlice := make([]string, len(inputSlice))

	for i, line := range inputSlice {
		maskedSlice[i] = s.maskingLinks(line)
	}

	err = s.pres.Present(maskedSlice)
	if err != nil {
		return fmt.Errorf("ошибка при записи файла: %v", err)
	}

	return nil
}
