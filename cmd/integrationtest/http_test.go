package integrationtest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/felixbecker/hexadiscountexample/api"
	internal "github.com/felixbecker/hexadiscountexample/internal/factory"
)

func Test_API_with_Redis_store(t *testing.T) {

	config := internal.Config{
		StoreType: "redis",
		Redis: internal.Redis{
			Addr: "localhost:6379",
		},
	}

	_ = config
	factory, err := internal.NewFactory()
	if err != nil {
		t.Errorf("This should not happen")
	}

	api := api.New(factory.Application())
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/discounter/3232323", nil)
	api.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected that status code to be: '%s'; got: '%s'", http.StatusText(http.StatusOK), http.StatusText(w.Code))
	}

	expected := `{"result":"2424242.000000"}`
	if w.Body.String() != expected {
		t.Errorf("expected the body to be: '%v';  got: '%v'", w.Body.String(), expected)
	}

}

func Test_API_with_Postgres_store(t *testing.T) {

	config := internal.Config{
		StoreType: "postgres",
		Postgres: internal.Postgres{
			User:      "root",
			Password:  "root",
			Host:      "discounter",
			Tablename: "rates",
		},
	}

	_ = config
	factory, err := internal.NewFactory()
	if err != nil {
		t.Errorf("This should not happen")
	}

	api := api.New(factory.Application())
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/discounter/3232323", nil)
	api.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected that status code to be: '%s'; got: '%s'", http.StatusText(http.StatusOK), http.StatusText(w.Code))
	}

	expected := `{"result":"2424242.000000"}`
	if w.Body.String() != expected {
		t.Errorf("expected the body to be: '%v';  got: '%v'", w.Body.String(), expected)
	}

}
func Test_API_with_in_Memory_store(t *testing.T) {

}

func Test_WebUI_with_Redis_store(t *testing.T) {

}
func Test_WebUI_with_Postgres_store(t *testing.T) {

}
func Test_WebUI_with_in_Memory_store(t *testing.T) {

}
