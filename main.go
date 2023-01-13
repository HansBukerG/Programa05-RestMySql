package main

import (
	"fmt"
	"github.com/HansBukerG/Programa05-RestMySql/app"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Programa 05 - Rest MySql para gestion de tabla")
	app.App_init()

}