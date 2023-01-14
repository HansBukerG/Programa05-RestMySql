package app_persona

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HansBukerG/Programa05-RestMySql/app/interfaces"
	"github.com/HansBukerG/Programa05-RestMySql/app/services/ioService"
	PersonaService "github.com/HansBukerG/Programa05-RestMySql/app/services/persona"
	"github.com/gorilla/mux"
)

var Persona interfaces.Persona

func CreatePersona(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Llamada a funcion CreatePersona()")
	//Se extrae la data desde la BD
	requestPersona := &Persona
	ioService.ParseBody(request, requestPersona)
	nuevoRegistro := PersonaService.CreatePersona(requestPersona)
	//Se prepara la data para ser insertada
	respuesta, _ := json.Marshal(nuevoRegistro)
	ioService.ExecuteResponse(writer, http.StatusAccepted, respuesta)
}
func GetAllPersonas(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Llamada a funcion GetAllPersonas()")
	RegistroPersonas := PersonaService.ReadPersonas()
	//Se prepara la data para ser insertada
	respuesta, _ := json.Marshal(RegistroPersonas)
	ioService.ExecuteResponse(writer, http.StatusAccepted, respuesta)
}
func GetPersonasById(writer http.ResponseWriter, request *http.Request) {
	variablesRequest := mux.Vars(request)
	idPersona := variablesRequest["personaId"]
	fmt.Println("llamada a funcion DeletePersonas con parámetro: "+idPersona )
	RegistroEncontrado := PersonaService.ReadPersonaById(idPersona)

	//Se prepara la data para ser insertada
	respuesta, _ := json.Marshal(RegistroEncontrado)
	ioService.ExecuteResponse(writer, http.StatusAccepted, respuesta)
}
func UpdatePersonas(writer http.ResponseWriter, request *http.Request) {
	variablesRequest := mux.Vars(request)
	id := variablesRequest["personaId"]
	fmt.Println("Llamada a funcion UpdaTePersonas con request id: " + id)

	personaRequest := &Persona
	ioService.ParseBody(request, personaRequest)

	PersonaService.UpdatePersonas(id,personaRequest)

	RegistroEncontrado := PersonaService.ReadPersonaById(id)

	//Se prepara la data para ser insertada
	respuesta, _ := json.Marshal(RegistroEncontrado)
	ioService.ExecuteResponse(writer, http.StatusAccepted, respuesta)

}
func DeletePersonas(writer http.ResponseWriter, request *http.Request) {
	
	variablesRequest := mux.Vars(request)
	idPersona := variablesRequest["personaId"]
	fmt.Println("llamada a funcion DeletePersonas con parámetro: "+idPersona )
	PersonaService.DeletePersonas(idPersona)

	personaVacia := &Persona

	//Se prepara la data para ser insertada
	respuesta, _ := json.Marshal(personaVacia)
	ioService.ExecuteResponse(writer, http.StatusAccepted, respuesta)
}
