package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	apiEcho "github.com/mateo-tavera/libraryPortal/apiEcho"
	apiGorilla "github.com/mateo-tavera/libraryPortal/apiGorilla"
)

func main() {

	//Using the command line to control the execution
	framework := string(os.Args[1])

	//Switch structure to select which framework execute
	switch framework {

	case "gorilla":
		//Create the rout pointer
		r := mux.NewRouter()

		//Create the api object
		a := &apiGorilla.API{}
		//Resgiter the routes to that api
		a.RegisterRoutes(r)

		//Start server
		fmt.Println("Using Gorilla/Mux")
		fmt.Println("Listening...")
		http.ListenAndServe(":8081", r)

	case "echo":
		//Create the rout pointer
		e := echo.New()

		//Create the api object
		a := &apiEcho.API{}

		//Resgiter the routes to that api
		a.RegisterRoutes(e)

		//Start server
		fmt.Println("Using Echo")
		e.Logger.Fatal(e.Start(":8081"))

	default:
		fmt.Println("No framework selected. App cannot start")
	}

}
