package apiecho

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/schema"
	"github.com/labstack/echo"
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
func (a *API) GetBooks(c echo.Context) error {
	//Add filter using query params
	BookParams := &BooksParams{}
	BookParams.To = len(BookList)
	err := SchemaDecoder.Decode(BookParams, c.QueryParams()) //Here we take the URL params and save them in the Struct BookParams

	//Validate all kind of errors
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)

	}

	if BookParams.From > len(BookList) || BookParams.From < 0 {
		return c.JSON(http.StatusBadRequest, err)

	}

	if BookParams.To < 0 || BookParams.To > len(BookList) {
		return c.JSON(http.StatusBadRequest, err)

	}

	//Output
	return c.JSON(http.StatusOK, BookList[BookParams.From:BookParams.To])

}

//Get single book using id
func (a *API) GetBook(c echo.Context) error {
	idBook, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)

	}

	if idBook-1 < 0 || idBook-1 > len(BookList)-1 {
		return c.JSON(http.StatusBadRequest, err)

	}
	//Output
	return c.JSON(http.StatusOK, BookList[idBook-1])
}

func (a *API) CreateBook(c echo.Context) error {
	var book = &Books{}

	err := json.NewDecoder(c.Request().Body).Decode(book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)

	}

	BookList = append(BookList, *book)
	return c.NoContent(http.StatusCreated)

}
