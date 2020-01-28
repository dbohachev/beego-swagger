package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dbohachev/beego-swagger/example/beego-api/models"
	_ "github.com/dbohachev/beego-swagger/example/beego-api/routers"
	"github.com/stretchr/testify/assert"

	"github.com/astaxie/beego"
)

func TestGetAll(t *testing.T) {
	models.Init()
	// Setup
	requestTests := []httpTestResult{
		// Default path to GET all users should succeed
		{
			RequestType: "GET",
			Path:        "/api/v1/users/",
			Expected:    http.StatusOK,
		},
	}

	// Run
	for _, test := range requestTests {
		r := httptest.NewRequest(test.RequestType, test.Path, nil)
		w := serveHttp(r)
		assert.Equal(t, test.Expected, w.Code)
	}
}

func TestGetSpecific(t *testing.T) {
	models.Init()
	// Setup
	requestTests := []httpTestResult{
		// User with this id does not exist, not found
		{
			RequestType: "GET",
			Path:        "/api/v1/users/100500",
			Expected:    http.StatusNotFound,
		},
		// This id is not valid, bad request
		{
			RequestType: "GET",
			Path:        "/api/v1/users/1a0500",
			Expected:    http.StatusBadRequest,
		},
		// Found user, success
		{
			RequestType: "GET",
			Path:        "/api/v1/users/2",
			Expected:    http.StatusOK,
		},
	}

	// Run
	for _, test := range requestTests {
		r := httptest.NewRequest(test.RequestType, test.Path, nil)
		w := serveHttp(r)
		assert.Equal(t, test.Expected, w.Code)
	}
}

func TestPost(t *testing.T) {
	models.Init()
	// Setup
	requestTests := []httpTestResult{
		// Post to "users"
		{
			RequestType: "POST",
			Path:        "/api/v1/users",
			Payload:     "",
			Expected:    http.StatusBadRequest,
		},
		// Success. Id is ignored.
		{
			RequestType: "POST",
			Path:        "/api/v1/users",
			Payload: `{
						"country": "Ukraine",
						"id": 1,
						"name": "Peter"
					}`,
			Expected: http.StatusOK,
		},
	}

	// Run
	for _, test := range requestTests {
		var body io.Reader = nil
		if len(test.Payload) != 0 {
			body = strings.NewReader(test.Payload)
		}
		r := httptest.NewRequest(test.RequestType, test.Path, body)
		w := serveHttp(r)
		assert.Equal(t, test.Expected, w.Code)
	}
}

func TestPut(t *testing.T) {
	models.Init()
	// Setup
	requestTests := []httpTestResult{
		// PUT to "users"
		{
			RequestType: "PUT",
			Path:        "/api/v1/users",
			Payload:     "",
			Expected:    http.StatusBadRequest,
		},
		// Success. User is updated
		{
			RequestType: "PUT",
			Path:        "/api/v1/users",
			Payload: `{
						"country": "USA",
						"id": 1,
						"name": "Dmytro"
					}`,
			Expected: http.StatusOK,
		},
		// Fail, no user by this ID
		{
			RequestType: "PUT",
			Path:        "/api/v1/users",
			Payload: `{
						"country": "USA",
						"id": 100500,
						"name": "Dmytro"
					}`,
			Expected: http.StatusNotFound,
		},
	}

	// Run
	for _, test := range requestTests {
		var body io.Reader = nil
		if len(test.Payload) != 0 {
			body = strings.NewReader(test.Payload)
		}
		r := httptest.NewRequest(test.RequestType, test.Path, body)
		w := serveHttp(r)
		assert.Equal(t, test.Expected, w.Code)
	}
}

func TestDelete(t *testing.T) {
	models.Init()
	// Setup
	requestTests := []httpTestResult{
		// DELETE to "users"
		{
			RequestType: "DELETE",
			Path:        "/api/v1/users",
			Expected:    http.StatusBadRequest,
		},
		// OK, success
		{
			RequestType: "DELETE",
			Path:        "/api/v1/users/1",
			Expected:    http.StatusOK,
		},
		// Not found
		{
			RequestType: "DELETE",
			Path:        "/api/v1/users/100500",
			Expected:    http.StatusNotFound,
		},
	}

	// Run
	for _, test := range requestTests {
		r := httptest.NewRequest(test.RequestType, test.Path, nil)
		w := serveHttp(r)
		assert.Equal(t, test.Expected, w.Code)
	}
}

// serveHttp - wrapper to start our requests to beego
func serveHttp(r *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	// proper json handling
	beego.BConfig.CopyRequestBody = true
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	return w
}

// wraps result of http execution
type httpTestResult struct {
	RequestType string
	Payload     string
	Path        string
	Expected    int
}
