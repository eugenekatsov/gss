package internal

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a new HTTP request with a JSON payload
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name":"John"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the CreateUser handler
	err := CreateUser(c)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert the response status code
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Assert the response body
	expected := `{"id":1,"name":"John"}`
	assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
}

func TestGetUser(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Add a user for testing
	users[1] = &user{
		ID:   1,
		Name: "John",
	}

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Call the GetUser handler
	err := GetUser(c)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Assert the response body
	expected := `{"id":1,"name":"John"}`
	assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
}

func TestUpdateUser(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Add a user for testing
	users[1] = &user{
		ID:   1,
		Name: "John",
	}

	// Create a new HTTP request with a JSON payload
	req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(`{"name":"Doe"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Call the UpdateUser handler
	err := UpdateUser(c)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Assert the response body
	expected := `{"id":1,"name":"Doe"}`
	assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))

	// Assert that the user has been updated
	assert.Equal(t, "Doe", users[1].Name)
}

func TestDeleteUser(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Add a user for testing
	users[1] = &user{
		ID:   1,
		Name: "John",
	}

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Call the DeleteUser handler
	err := DeleteUser(c)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert the response status code
	assert.Equal(t, http.StatusNoContent, rec.Code)

	// Assert that the user has been deleted
	_, exists := users[1]
	assert.False(t, exists)
}

func TestHealthz(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the Healthz handler
	err := Healthz(c)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Assert the response body
	assert.Equal(t, "200", strings.TrimSpace(rec.Body.String()))
}
