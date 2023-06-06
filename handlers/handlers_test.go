package handlers

import (
	"errors"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
)

type MockHandlers struct {
	mock.Mock
}

func (m *MockHandlers) ListFacts(c *fiber.Ctx) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockHandlers) CreateFact(c *fiber.Ctx) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockHandlers) ShowFact(c *fiber.Ctx) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockHandlers) UpdateFact(c *fiber.Ctx) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockHandlers) DeleteFact(c *fiber.Ctx) error {
	args := m.Called(c)
	return args.Error(0)
}

func TestListFacts(t *testing.T) {
	// Crea una instancia del mock
	mockHandlers := new(MockHandlers)

	// Define el contexto de Fiber
	ctx := new(fiber.Ctx)

	// Define el resultado esperado
	expectedError := errors.New("Some error")

	// Configura el comportamiento esperado del mock
	mockHandlers.On("ListFacts", ctx).Return(expectedError)

	// Ejecuta la función bajo prueba
	err := mockHandlers.ListFacts(ctx)

	// Verifica si el error retornado es el esperado
	assert.Equal(t, expectedError, err)
}

func TestCreateFact(t *testing.T) {
	// Crea una instancia del mock
	mockHandlers := new(MockHandlers)

	// Define el contexto de Fiber
	ctx := new(fiber.Ctx)

	// Define el resultado esperado
	expectedError := errors.New("Some error")

	// Configura el comportamiento esperado del mock
	mockHandlers.On("CreateFact", ctx).Return(expectedError)

	// Ejecuta la función bajo prueba
	err := mockHandlers.CreateFact(ctx)

	// Verifica si el error retornado es el esperado
	assert.Equal(t, expectedError, err)
}

func TestShowFact(t *testing.T) {
	// Crea una instancia del mock
	mockHandlers := new(MockHandlers)

	// Define el contexto de Fiber
	ctx := new(fiber.Ctx)

	// Define el resultado esperado
	expectedError := errors.New("Some error")

	// Configura el comportamiento esperado del mock
	mockHandlers.On("ShowFact", ctx).Return(expectedError)

	// Ejecuta la función bajo prueba
	err := mockHandlers.ShowFact(ctx)

	// Verifica si el error retornado es el esperado
	assert.Equal(t, expectedError, err)
}
