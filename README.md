# libraryPortal
This app uses different frameworks to implement REST APIs in a simulated library

The core of the project is to determine the general aspects, advantages, implementation, resources, convenience and the facility to change between Golang frameworks when creating APIs. All this is performed under the context of a library with simple CRUD operations.

The app is executied using the command line prompt, writing the framework after the .exe file:

first you get the binary file:

`go build main.go`

then you execute ir adding the framework, e.g:

`./ main.exe gorilla`

## The frameworks or libraries used:

* [gorilla/mux](https://github.com/mateo-tavera/libraryPortal/tree/main/apiGorilla)-> `gorilla`
* [echo](https://github.com/mateo-tavera/libraryPortal/tree/main/apiEcho)-> `echo`
