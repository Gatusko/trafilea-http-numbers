package controller

import (
	"fmt"
	"github.com/Gatusko/trafilea-http-numbers/domain/model"
	"github.com/Gatusko/trafilea-http-numbers/services"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type NumberController struct {
	numberService services.NumberService
}

func NewNumberController(numberService services.NumberService) *NumberController {
	return &NumberController{
		numberService: numberService,
	}
}

func (n *NumberController) postNumber(w http.ResponseWriter, r *http.Request) {
	number := model.NewNumber()
	err := readJson(r, &number)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = number.Validate()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	val, err := n.numberService.AddNumber(number)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, val)
}

func (n *NumberController) getNumbers(w http.ResponseWriter, r *http.Request) {
	vals, err := n.numberService.GetAllNumbers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, vals)
}

func (n *NumberController) deleteNumber(w http.ResponseWriter, r *http.Request) {
	num := chi.URLParam(r, "number")
	if num == "" {
		respondWithError(w, http.StatusBadRequest, "Need to provided Id of number for delete")
		return
	}
	numInt, err := strconv.Atoi(num)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Need to provided a Correct Number")
		return
	}
	err = n.numberService.DeleteNumber(numInt)
	if err != nil {
		switch {
		case fmt.Sprintf("Value not found: %v", numInt) == err.Error():
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	respondWithoutBody(w, http.StatusOK)
}

func (n *NumberController) getNumber(w http.ResponseWriter, r *http.Request) {
	num := chi.URLParam(r, "number")
	if num == "" {
		respondWithError(w, http.StatusBadRequest, "Need to provided Id of number for delete")
		return
	}
	numInt, err := strconv.Atoi(num)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Need to provided a Correct Number")
		return
	}
	val, err := n.numberService.GetNumber(numInt)
	if err != nil {
		switch {
		case fmt.Sprintf("Value not found: %v", numInt) == err.Error():
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	respondWithJson(w, http.StatusOK, val)
}

// Here we define the routes and  see wich one can use middleware or not
func (n *NumberController) NumberRoutes() http.Handler {
	route := chi.NewRouter()
	route.Use(logRoute)
	route.Post("/", n.postNumber)
	route.Get("/", n.getNumbers)
	route.Delete("/{number}", n.deleteNumber)
	route.Get("/{number}", n.getNumber)
	return route
}
