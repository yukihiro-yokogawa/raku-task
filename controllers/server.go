package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartMainServer() error {
	log.Println("Rest Api With Mux Routers")
	router := mux.NewRouter().StrictSlash(true)
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
