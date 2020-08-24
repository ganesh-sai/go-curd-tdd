package tdd

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// App
type App struct {
	Router *http.ServeMux
	DB     *sql.DB
}

// Initialize - initialzie the applicaiton with db connection
func (app *App) Initialize(username, password, dbname string) {
	connection := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, dbname)
	var err error
	app.DB, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatalln(err)
	}

	app.Router = &http.ServeMux{}

	app.InitializeRoutes()
}

// InitializeRoutes - register routes
func (app *App) InitializeRoutes() {
	app.Router.HandleFunc("/", helloworld)
}

func helloworld(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello World..!"))
}

// Run - Starts the server
func (app *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, app.Router))
}
