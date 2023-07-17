package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go-m.-azwan-alterra/practikum13/lib/database"
	"go-m.-azwan-alterra/practikum13/model"

	"github.com/labstack/echo/v4"
)

func TestGetUserController(t *testing.T) {
	e := echo.New()

	// mock database
	users := []model.User{
		{
			ID:   1,
			Name: "Azwan",
		},
	}
	database.GetUsers = func() ([]model.User, error) {
		return users, nil
	}

	// test GetUserController
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := GetUserController(c)
	if err != nil {
		t.Errorf("GetUserController error: %v", err)
	}

	// assert response
	if rec.Code != http.StatusOK {
		t.Errorf("GetUserController status code: %d", rec.Code)
	}

	var response map[string]interface{}
	err = rec.Body.Unmarshal(&response)
	if err != nil {
		t.Errorf("Unmarshal response error: %v", err)
	}

	if response["status"] != "success" {
		t.Errorf("response status: %s", response["status"])
	}

	if len(response["users"].([]model.User)) != 1 {
		t.Errorf("response users length: %d", len(response["users"].([]model.User)))
	}
}

func TestCreateUserController(t *testing.T) {
	e := echo.New()

	// mock database
	user := model.User{
		Name: "Azwan",
	}
	database.CreateUser = func(u *model.User) error {
		*u = user
		return nil
	}

	// test CreateUserController
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := CreateUserController(c)
	if err != nil {
		t.Errorf("CreateUserController error: %v", err)
	}

	// assert response
	if rec.Code != http.StatusCreated {
		t.Errorf("CreateUserController status code: %d", rec.Code)
	}

	var response map[string]interface{}
	err = rec.Body.Unmarshal(&response)
	if err != nil {
		t.Errorf("Unmarshal response error: %v", err)
	}

	if response["status"] != "success" {
		t.Errorf("response status: %s", response["status"])
	}

	if response["user"].(model.User) != user {
		t.Errorf("response user: %v", response["user"])
	}
}
