package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/teguhbudhi13/users/app/handler"
	"github.com/teguhbudhi13/users/app/model"
	"github.com/teguhbudhi13/users/config"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// routing for handling the users
	user := a.Router.PathPrefix("/users").Subrouter()
	user.HandleFunc("", a.handleRequest(handler.CreateUser)).Methods("POST")
	user.HandleFunc("", a.handleRequest(handler.GetAllUsers)).Methods("GET")
	user.HandleFunc("/{id:[0-9]+}", a.handleRequest(handler.GetUser)).Methods("GET")
	user.HandleFunc("/{id:[0-9]+}", a.handleRequest(handler.DeleteUser)).Methods("DELETE")
	user.HandleFunc("/{id:[0-9]+}", a.handleRequest(handler.UpdateUser)).Methods("PUT")
}

// Run router
func (a *App) Run(host string, allowedOrigins string) {
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{allowedOrigins})
	log.Fatal(http.ListenAndServe(host, handlers.CORS(headers, methods, origins)(a.Router)))
}

// RequestHandlerFunction request handler
type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
