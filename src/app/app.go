package app

import (
	"net/http"
	"handler"
	"config"
	"github.com/jinzhu/gorm"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	gormtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/jinzhu/gorm"
	"github.com/go-sql-driver/mysql"
)

type App struct {
	DB     *gorm.DB
}

var mux = muxtrace.NewRouter(muxtrace.WithServiceName("my_service"))

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	sqltrace.Register("mysql", &mysql.MySQLDriver{}, sqltrace.WithServiceName("my_service"))

	var err error = nil
	a.DB, err = gormtrace.Open(config.DB.Dialect, dbURI)

	if err != nil {
		log.Fatal("Could not connect database "+ err.Error())
	}

	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/", a.GetAllTest)
	a.Get("/lele", a.GetAllTest2)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	http.ListenAndServe(host,mux)
}

func (a *App) GetAllTest(w http.ResponseWriter, r *http.Request) {
	handler.GetAllTest(a.DB, w, r)
}
func (a *App) GetAllTest2(w http.ResponseWriter, r *http.Request) {
	handler.GetAllTest2(a.DB, w, r)
}
