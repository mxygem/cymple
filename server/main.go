package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/tidwall/pretty"
)

const (
	homepageTmpl = "./server/templates/homepage.html"
)

type employee struct {
	ID     int64  `json:"id,omitempty"`
	Gender string `json:"gender"`
}

type API struct {
	employees  []*employee
	employeesM []byte
}

func main() {
	var port string

	flag.StringVar(&port, "port", "3000", "port to listen to")
	flag.Parse()

	employees := []*employee{
		{ID: 1, Gender: "Male"},
		{ID: 2, Gender: "Female"},
		{ID: 3, Gender: "Nonbinary"},
		{ID: 4, Gender: "Female"},
	}

	allEmployees, err := json.Marshal(employees)
	if err != nil {
		log.Fatal(err)
	}

	api := &API{
		employees:  employees,
		employeesM: allEmployees,
	}

	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/", api.home)

		r.Route("/employees", func(r chi.Router) {
			r.Get("/", api.getAll)
			r.Get("/{id}", api.getEmployee)
		})
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go log.Print("employees api started on port " + port)
	go func() {
		if err := http.ListenAndServe(":"+port, r); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()
	<-stop

	fmt.Println("shutting down server")
}

func (a *API) home(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(homepageTmpl))

	err := t.Execute(rw, a.employees)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/html")
}

func (a *API) getAll(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(pretty.Pretty(a.employeesM)) // nolint: errcheck
}

func (a *API) getEmployee(rw http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, e := range a.employees {
		if id == int(e.ID) {
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(pretty.Pretty([]byte(fmt.Sprintf(`{"id":%d,"gender":"%s"}`, e.ID, e.Gender)))) // nolint: errcheck
			return
		}
	}

	rw.WriteHeader(http.StatusNotFound)
}
