package handlers

import (
	"golang-fifa-world-cup-web-service/data"
	"net/http"
)

// RootHandler returns an empty body status code
func RootHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNoContent)
}

// ListWinners returns winners from the list
func ListWinners(res http.ResponseWriter, req *http.Request) {
	year := req.URL.Query().Get("year")
	res.Header().Add("Content-Type", "application/json")
	if year == "" {
		winners, err := data.ListAllJSON()
		if err != nil {
			writeServerError(res, err)
			return
		}
		res.WriteHeader(http.StatusOK)
		res.Write(winners)
	} else {
		winners, err := data.ListAllByYear(year)
		if err != nil {
			writeServerError(res, err)
			return
		}
		res.Write(winners)
	}
}

func writeServerError(res http.ResponseWriter, err error) {
	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte(err.Error()))
}

// AddNewWinner adds new winner to the list
func AddNewWinner(res http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("X-ACCESS-TOKEN")
	isTokenValid := data.IsAccessTokenValid(accessToken)
	if !isTokenValid {
		res.WriteHeader(http.StatusUnauthorized)
		return
	} else {
		err := data.AddNewWinner(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		res.WriteHeader(http.StatusCreated)
		return
	}

}

// WinnersHandler is the dispatcher for all /winners URL
func WinnersHandler(res http.ResponseWriter, req *http.Request) {

}
