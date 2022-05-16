package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Application interface {
	Discount(amount float32) float32
}

type Api struct {
	app    Application
	router *mux.Router
}

func New(application Application) *Api {

	api := Api{}
	api.app = application
	api.routes()
	return &api
}

func (a *Api) routes() {
	a.router = mux.NewRouter()
	a.router.HandleFunc("/discounter/{amount}", a.disountHandler).Methods("POST")
	a.router.HandleFunc("/", a.index).Methods("GET")
}

func (a *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

func (a *Api) disountHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	amountString := vars["amount"]
	if amountString == "" {
		http.Error(w, `{"error":"not a valid amount"}`, http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseFloat(amountString, 32)
	if err != nil {
		http.Error(w, `{"error":"not a valid amount"}`, http.StatusBadRequest)
		return
	}

	result := a.app.Discount(float32(amount))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"result":"%f"}`, result)

}
func (a *Api) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"link":"http://%s/api/discounter"}`, r.Host)
}
