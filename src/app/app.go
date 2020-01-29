package app

import (
	"net/http"

	"github.com/Reckner/rad_product/src/controllers"
	"github.com/Reckner/rad_product/src/database"
	handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// InitService launches service
func InitService() {

	database.InitializeDB()

	router := mux.NewRouter().StrictSlash(true)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/items", controllers.GetItems).Methods("GET")

	if err := http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)); err != nil {
		panic(err)
	}

}
