package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HansBukerG/Programa05-RestMySql/app/httpHandlers"
	"github.com/HansBukerG/Programa05-RestMySql/app/routes"
	"github.com/gorilla/mux"
)

func App_init() {
	fmt.Println("Conexion con servidor mysql");
	HttpHandlers.ConexionMysql();

	route := mux.NewRouter();

	routes.RegisterPersonaRoutes(route)
	http.Handle("/",route)

	log.Fatal(http.ListenAndServe("localhost:8000", route))
}
