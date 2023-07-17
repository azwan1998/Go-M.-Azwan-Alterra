package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go-m.-azwan-alterra/practikum13/config"
	"go-m.-azwan-alterra/practikum13/controller"
	"go-m.-azwan-alterra/practikum13/model"
	"gorm.io/gorm"
)

func TestGetBooks(t *testing.T) {
	// Mock data
	mockBooks := []model.Buku{
		{Judul: "Book 1", Pengarang: "Author 1"},
		{Judul: "Book 2", Pengarang: "Author 2"},
		// Add more books if needed for your test cases
	}

	// Mock the database.GetBuku() function
	database.buku = func() ([]model.Buku, error) {
		return mockBooks, nil
	}

	// Create a new Echo context for the test
	e := echo.New()

	// Perform the HTTP request and record the response
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the handler function
	err := controller.GetBooks(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response["status"])
	assert.NotNil(t, response["users"])

	// Add more assertions based on your response data structure if needed
}

func TestCreateBooks(t *testing.T) {
	// Create a new Echo context for the test
	e := echo.New()

	// Mock the request payload
	requestBody := []byte(`{"title": "New Book", "author": "New Author"}`)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the handler function
	err := controller.CreateBooks(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var createdBook model.Buku
	err = json.Unmarshal(rec.Body.Bytes(), &createdBook)
	assert.NoError(t, err)
	assert.Equal(t, "New Book", createdBook.Judul)
	assert.Equal(t, "New Author", createdBook.Pengarang)

	// You can also assert the book was created in the database if possible
}

func TestShow(t *testing.T) {
	// Mock data
	mockBook := model.Buku{Pengarang: "Mock Book", Penerbit: "Mock Author"}

	// Mock the database query function (config.DB.First)
	config.InitDB.First = func(out interface{}, where ...interface{}) *gorm.DB {
		b, ok := out.(*model.Buku)
		assert.True(t, ok, "should pass a pointer to model.Buku")
		assert.Equal(t, 1, where[0], "should query with ID 1")

		// Set the mock book data to the output
		*b = mockBook

		// Return an empty Gorm DB object to mock success
		return &gorm.DB{}
	}

	// Call the Show function
	book, err := controller.Show("1")

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, book)
	assert.Equal(t, &mockBook, book)
}
