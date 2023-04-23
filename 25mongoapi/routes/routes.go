package routes

import (
	"25mongoapi/controllers"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	controllers.Init()

	router.HandleFunc("/api/movies", controllers.InsertOneMovie).Methods("POST")
	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies", controllers.DeleteAllMovies).Methods("DELETE")
	router.HandleFunc("/api/movies/{movieId}", controllers.UpdateOneMovie).Methods("PUT")
	router.HandleFunc("/api/movies/{movieId}", controllers.DeleteOneMovie).Methods("DELETE")

	return router
}
