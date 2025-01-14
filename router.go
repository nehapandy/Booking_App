package Router

import (
	"BOOKING-APP/Router/controller"
	"fmt"

	"github.com/gorilla/mux"
)

func Initialize() {
	fmt.Println("Router Initialized")
}

// Router initializes the routes and returns the router instance.
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/products", controller.GetMyAllProducts).Methods("GET")
	router.HandleFunc("/api/product", controller.CreateProduct).Methods("POST")
	router.HandleFunc("/api/product/{id}", controller.MarkAsBuy).Methods("PUT")
	router.HandleFunc("/api/product/{id}", controller.DeleteAProduct).Methods("DELETE")
	router.HandleFunc("/api/deleteallproducts", controller.DeleteALLProducts).Methods("DELETE")

	return router
}
