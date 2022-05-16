package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Application is an interface that acts as a port
type Application interface {
	Discount(amount float32) float32
}

type Web struct {
	app    Application
	router *mux.Router
}

func New(application Application) *Web {
	w := Web{}
	w.app = application
	w.routes()
	return &w
}

func (web *Web) routes() {
	web.router = mux.NewRouter()

	web.router.HandleFunc("/", web.homePage)
	web.router.HandleFunc("/discounter", web.discountViewGet).Methods("GET")
	web.router.HandleFunc("/discounter", web.discountViewPost).Methods("POST")

}

const indexHTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Calculate a discount - index</title>
</head>
<body>
  <a href="/discounter">Please calculate your favorite discount</a>
</body>
</html>
`

func (web *Web) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, indexHTML)
}

const inputForm = `<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Calculate a discount - index</title>
</head>

<body>
  <h1>Disount calculation</h1>
  <form action="/discounter" method="post">
    <label for="amount">Please enter a number to calculate your discount:</label>
    <input type="number" id="amount" name="amount"/>
    <button type="submit">Calculate</button>
  </form>
</body>

</html>`

func (web *Web) discountViewGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, inputForm)
}

const resultPage = `<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Calculate a discount - index</title>
</head>

<body>
  <h1>Disount calculation</h1>
	<p>based on your amount: {{ .Amount}} the provided discount is: {{.Result}}</p>
	<p>start a new calculation <a href="/discounter">here</a> or go straight to <a href="/">home</a></p>
</body>

</html>`

func (web *Web) discountViewPost(w http.ResponseWriter, r *http.Request) {
	resultTempl, _ := template.New("result").Parse(resultPage)
	errorTmpl, _ := template.New("error").Parse(errorPage)

	r.ParseForm()
	amountString := r.FormValue("amount")
	if amountString == "" {

		log.Println("Amount string is empty")
		w.WriteHeader(http.StatusBadRequest)
		errorTmpl.Execute(w, "Amount cannot be empty")

		return
	}

	amount, err := strconv.ParseFloat(amountString, 32)

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusBadRequest)
		errorTmpl.Execute(w, err)
		return
	}

	result := web.app.Discount(float32(amount))
	resultTempl.Execute(w, struct {
		Amount float32
		Result float32
	}{
		Amount: float32(amount),
		Result: result,
	})
}

var errorPage = `<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Calculate a discount - index</title>
</head>

<body>
  <h1>Error: please start a new calculation</h1>
	<p>{{ . }}</p>
	<a href="/">start a new Calculation here</a>
</body>

</html>`

func (web *Web) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	web.router.ServeHTTP(w, r)
}
