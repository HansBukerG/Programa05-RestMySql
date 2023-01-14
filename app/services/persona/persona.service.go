package PersonaService

import (
	"strconv"

	"github.com/HansBukerG/Programa05-RestMySql/app/interfaces"
	HttpService "github.com/HansBukerG/Programa05-RestMySql/app/services/httpService"
)

func CreatePersona(Persona *interfaces.Persona) interfaces.Persona {

	query := "insert into Personas2(rut,nombre,apellido) Values('" + Persona.Rut + "','" + Persona.Nombre + "','" + Persona.Apellido + "')"
	idInsert, _ := HttpService.EjecucionConsulta(query)
	nuevaQuery := "Select id,rut,nombre,apellido from personas2 where id = " + strconv.FormatInt(idInsert, 10)
	NuevoRegistro := HttpService.ExtraerRegistro(nuevaQuery)
	NuevaPersona := interfaces.Persona{}
	for NuevoRegistro.Next() {
		var id int
		var rut, nombre, apellido string
		err := NuevoRegistro.Scan(&id, &rut, &nombre, &apellido)
		if err != nil {
			panic(err.Error())
		}
		NuevaPersona.Id = strconv.Itoa(id)
		NuevaPersona.Rut = rut
		NuevaPersona.Nombre = nombre
		NuevaPersona.Apellido = apellido
	}
	return NuevaPersona
}

func ReadPersonas() []interfaces.Persona {
	query := "Select id, rut, nombre,apellido from personas2"
	persona := interfaces.Persona{}
	arregloPersonas := []interfaces.Persona{}

	registrosEncontrados := HttpService.ExtraerRegistro(query)

	for registrosEncontrados.Next() {
		var id int
		var rut, nombre, apellido string
		err := registrosEncontrados.Scan(&id, &rut, &nombre, &apellido)
		if err != nil {
			panic("Error en la transformación de registros")
		}
		persona.Id = strconv.Itoa(id)
		persona.Rut = rut
		persona.Nombre = nombre
		persona.Apellido = apellido

		arregloPersonas = append(arregloPersonas, persona)
	}
	return arregloPersonas
}

func ReadPersonaById(id string) interfaces.Persona {
	query := "select id,rut,nombre,apellido from personas2 where id =" + id
	persona := interfaces.Persona{}

	registroEncontrado := HttpService.ExtraerRegistro(query)

	for registroEncontrado.Next() {
		var id int
		var rut, nombre, apellido string
		registroEncontrado.Scan(&id, &rut, &nombre, &apellido)
		persona.Id = strconv.Itoa(id)
		persona.Rut = rut
		persona.Nombre = nombre
		persona.Apellido = apellido
	}
	return persona
}

func UpdatePersonas(id string, persona *interfaces.Persona) {
	query := "UPDATE personas2 SET rut = '"+persona.Rut+"', nombre = '"+persona.Nombre+"', apellido = '"+persona.Apellido+"' WHERE personas2.id =" + id
	HttpService.EjecucionConsultaVoid(query) 
}

func DeletePersonas(id string) {
	query := "delete from personas2 where id = " + id

	HttpService.EjecucionConsultaVoid(query)
}
