package routes

import (
	"imoveis-back/controllers"

	"github.com/gorilla/mux"
)

func StartRoutes(r *mux.Router) {
	r.HandleFunc("/apartamentos", controllers.GetAllApartaments).Methods("GET")
	r.HandleFunc("/apartamentos/{id}", controllers.GetApartamentById).Methods("GET")
	r.HandleFunc("/apartamentos", controllers.InsertApartament).Methods("POST")
	r.HandleFunc("/apartamentos/{id}", controllers.UpdateApartament).Methods("PUT")
	r.HandleFunc("/apartamentos/{id}", controllers.DeleteApartament).Methods("DELETE")

	r.HandleFunc("/casas", controllers.GetAllHouses).Methods("GET")
	r.HandleFunc("/casas/{id}", controllers.GetHouseById).Methods("GET")
	r.HandleFunc("/casas", controllers.InsertHouse).Methods("POST")
	r.HandleFunc("/casas/{id}", controllers.UpdateHouse).Methods("PUT")
	r.HandleFunc("/casas/{id}", controllers.DeleteHouse).Methods("DELETE")

	r.HandleFunc("/terrenos", controllers.GetAllLands).Methods("GET")
	r.HandleFunc("/terrenos/{id}", controllers.GetAllLands).Methods("GET")
	r.HandleFunc("/terrenos", controllers.InsertLand).Methods("POST")
	r.HandleFunc("/terrenos/{id}", controllers.UpdateLand).Methods("PUT")
	r.HandleFunc("/terrenos/{id}", controllers.DeleteLand).Methods("DELETE")
}
