package routes

import (
	"github.com/HansBukerG/Programa05-RestMySql/app/components/persona"
	"github.com/gorilla/mux"
)

var RegisterPersonaRoutes = func(router *mux.Router){
	router.HandleFunc("/persona/", app_persona.CreatePersona ).Methods("POST")
	router.HandleFunc("/persona/", app_persona.GetAllPersonas ).Methods("GET")
	router.HandleFunc("/persona/{personaId}", app_persona.GetPersonasById  ).Methods("GET")
	router.HandleFunc("/persona/{personaId}", app_persona.UpdatePersonas ).Methods("PUT")
	router.HandleFunc("/persona/{personaId}", app_persona.DeletePersonas ).Methods("DELETE")
}