package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockProducer struct {
	mock.Mock
}

func (m *MockProducer) Produce() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

type MockPresenter struct {
	mock.Mock
}

func (m *MockPresenter) Present(data []string) error {
	args := m.Called(data)
	return args.Error(0)
}

func TestService_Run_Success(t *testing.T) {
	mockProducer := new(MockProducer)
	mockPresenter := new(MockPresenter)

	mockProducer.On("Produce").Return([]string{
		"Visit my page: http://example.com",
	}, nil)

	mockPresenter.On("Present", []string{
		"Visit my page: http://***********",
	}).Return(nil)

	service := NewService(mockProducer, mockPresenter)

	err := service.Run()

	assert.Nil(t, err)
}

func TestService_Run_ProduceError(t *testing.T) {
	mockProducer := new(MockProducer)
	mockPresenter := new(MockPresenter)

	mockProducer.On("Produce").Return([]string{}, errors.New("produceError"))

	mockPresenter.On("Present", []string{
		"Visit my page: http://***********",
	}).Return(nil)

	service := Service{prod: mockProducer, pres: mockPresenter}

	err := service.Run()

	assert.EqualError(t, err, "ошибка при чтении файла: produceError")
}

func TestService_Run_PresentError(t *testing.T) {
	mockProducer := new(MockProducer)
	mockPresenter := new(MockPresenter)

	mockProducer.On("Produce").Return([]string{
		"Visit my page: http://example.com",
	}, nil)

	mockPresenter.On("Present", []string{
		"Visit my page: http://***********",
	}).Return(errors.New("presentError"))

	service := NewService(mockProducer, mockPresenter)

	err := service.Run()

	assert.EqualError(t, err, "ошибка при записи файла: presentError")
}
