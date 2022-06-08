package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	apiEcho "github.com/mateo-tavera/libraryPortal/apiEcho"
	apiGorilla "github.com/mateo-tavera/libraryPortal/apiGorilla"
)

func main() {

	framework := string(os.Args[1])

	switch framework {

	case "gorilla":
		//Create the rout pointer
		r := mux.NewRouter()

		//Create the api object
		a := &apiGorilla.API{}
		//Resgiter the routes to that api
		a.RegisterRoutes(r)

		//Start server
		fmt.Println("Listening...")
		http.ListenAndServe(":8081", r)
	case "echo":
		e := echo.New()
		a := &apiEcho.API{}
		a.RegisterRoutes(e)
		e.Logger.Fatal(e.Start(":8081"))
	default:
		fmt.Println("No framework selected. App cannot start")
	}

}
