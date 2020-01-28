package beegoSwagger

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"
	_ "github.com/dbohachev/beego-swagger/example/beego-api/docs" // Needed for json testing
	"github.com/stretchr/testify/assert"
)

func Init() {
	// Set handlers on package init.
	beego.Get("/swagger/*", Handler())
}

func TestHandler(t *testing.T) {
	Init()
	// Setup
	requestTests := []struct {
		RequestType string
		Path        string
		Expected    int
	}{
		// Default (empty) path should be "Moved Permanently"
		{
			RequestType: "GET",
			Path:        "/swagger/",
			Expected:    http.StatusMovedPermanently,
		},
		{
			RequestType: "GET",
			Path:        "/swagger/index.html",
			Expected:    http.StatusOK,
		},
		{
			RequestType: "GET",
			Path:        "/swagger/doc.json",
			Expected:    http.StatusOK,
		},
		{
			RequestType: "GET",
			Path:        "/swagger/favicon-16x16.png",
			Expected:    http.StatusOK,
		},
		{
			RequestType: "GET",
			Path:        "/swagger/notfound",
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
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	return w
}
