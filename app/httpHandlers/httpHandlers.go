package HttpHandlers

import (
	"database/sql"
	"fmt"
)
var ConexionAbierta *sql.DB
/*
	Se establece una conexión a MySQL con los siguientes campos:
		driver := "mysql"
		user := "root"
		password := ""
		db_name := "curso_golang"
*/
func ConexionMysql() {
	driver := "mysql"
	user := "root"
	password := ""
	db_name := "curso_golang"

	conexion, err := sql.Open(driver, user+":"+password+"@tcp(127.0.0.1)/"+db_name)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Conexion establecida con exito!")
		ConexionAbierta = conexion
	}
}
/*
Esta función ejecuta un query entrante
*/
func EjecucionConsulta(query string) {
	// ConexionAbierta.Ping();
	insertarRegistro, err := ConexionAbierta.Prepare(query)

	if err != nil {
		panic(err.Error())
	} else {
		response, responseError := insertarRegistro.Exec()
		if responseError != nil {
			panic("Ha ocurrido un error: " + responseError.Error())
		} else {
			fmt.Println("El comando: ")
			fmt.Println("	" + query)
			fmt.Println("Ha sido ejecutado con exito!")
			fmt.Println("Respuesta:")
			fmt.Println(response)
		}
	}
}
/*
Esta función retorna un array de registros en base a la query que se este solicitando
*/
func ExtraerRegistro(query string) *sql.Rows {

	registrosEncontrados, err := ConexionAbierta.Query(query)
	if err != nil {
		panic("Fallo en la ejecución de comando: " + query + " Error: " + err.Error())
	}
	return registrosEncontrados
}