package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestSetupRoutes(t *testing.T) {
	app := fiber.New()

	// Create instance of the mock
	mockHandlers := new(MockHandlers)

	// Configurar expectativas de los mocks
	mockHandlers.On("ListFacts", mock.Anything).Return(nil).Once()
	mockHandlers.On("CreateFact", mock.Anything).Return(nil).Once()
	mockHandlers.On("ShowFact", mock.Anything).Return(nil).Once()
	mockHandlers.On("UpdateFact", mock.Anything).Return(nil).Once()
	mockHandlers.On("DeleteFact", mock.Anything).Return(nil).Once()

	// Ejecutar la función a probar
	SetupRoutes(app, mockHandlers)

	// Obtener las rutas registradas en el enrutador de Fiber
	routes := app.GetRoutes()

	// Verificar que se hayan registrado las rutas esperadas
	assert.Equal(t, 5, len(routes))

	req := httptest.NewRequest("GET", "/fact", nil)
	_ , err := app.Test(req)
	req = httptest.NewRequest("POST", "/fact", nil)
	_ , err = app.Test(req)
	req = httptest.NewRequest("GET", "/fact/1", nil)
	_ , err = app.Test(req)
	req = httptest.NewRequest("PATCH", "/fact/1", nil)
	_ , err = app.Test(req)
	req = httptest.NewRequest("DELETE", "/fact/1", nil)
	_ , err = app.Test(req)
	
	mockHandlers.AssertCalled(t, "ListFacts", mock.Anything)
	mockHandlers.AssertCalled(t, "CreateFact", mock.Anything)
	mockHandlers.AssertCalled(t, "ShowFact", mock.Anything)
	mockHandlers.AssertCalled(t, "UpdateFact", mock.Anything)
	mockHandlers.AssertCalled(t, "DeleteFact", mock.Anything)

	assert.NoError(t, err)

	// Verificar que se haya llamado a los métodos esperados en los mocks
	mockHandlers.AssertExpectations(t)
}

