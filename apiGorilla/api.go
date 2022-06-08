package apigorilla

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type API struct{}

type Books struct {
	Title string `json:"title"`
	Isbn  string `json:"isbn"`
}

//Schema is used for taking the params provided in the URL and use them as struct attribute
type BooksParams struct {
	From int `schema:"from"`
	To   int `schema:"to"`
}

var BookList = []Books{{"Mobey Dick", "1"}, {"Dracula", "2"}, {"Oliver Twist", "3"}, {"Frankenstein", "4"}, {"Great Expectations", "5"}}
var SchemaDecoder = schema.NewDecoder() //Create schema decoder

//Get all books according to limits provided un query params
func (a *API) GetBooks(w http.ResponseWriter, r *http.Request) {
	//Add filter using query params
	BookParams := &BooksParams{}
	BookParams.To = len(BookList)
	err := SchemaDecoder.Decode(BookParams, r.URL.Query()) //Here we take the URL params and save them in the Struct BookParams
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if BookParams.From > len(BookList) || BookParams.From < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if BookParams.To < 0 || BookParams.To > len(BookList) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(BookList[BookParams.From:BookParams.To])
}

//Get single book using id
func (a *API) GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idBook, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if idBook-1 < 0 || idBook-1 > len(BookList)-1 {
		w.WriteHeader(http.StatusNotFound)
		return

	}

	json.NewEncoder(w).Encode(BookList[idBook-1])
}

func (a *API) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book = &Books{}

	err := json.NewDecoder(r.Body).Decode(book)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	BookList = append(BookList, *book)
	w.WriteHeader(http.StatusCreated)

}
