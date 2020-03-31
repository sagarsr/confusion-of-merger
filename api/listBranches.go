package api

import (
	"bankapp/models"
	"bankapp/utils"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//ListOfBranches is the response format for viewing all branches
type ListOfBranches struct {
	TotalCount int             `json:"total_count"`
	Data       []models.Branch `json:"data"`
}

//ViewBranchList is to list all branches given bankname and city
var ViewBranchList = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var branchList models.Branch
	bankNameList, queryErr := r.URL.Query()["bank_name"]
	if !queryErr {
		utils.RespondWithError(w, http.StatusBadRequest, "Bank name to be given")
		return

	}
	bankName := bankNameList[0]
	cityNameList, queryErr := r.URL.Query()["city"]
	if !queryErr {
		utils.RespondWithError(w, http.StatusBadRequest, "city name to be given")
		return

	}
	cityName := cityNameList[0]
	limit := 10
	offset := 10
	switch r.Method {
	case "GET":
		log.Info(bankName, cityName, limit, offset)
		listofbranches := ListOfBranches{}
		branchlist, count, err := branchList.FetchBranches(utils.Connection, bankName, cityName, limit, offset)

		log.Info(branchlist)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
			return
		}
		listofbranches.TotalCount = count
		listofbranches.Data = branchlist
		utils.RespondWithJSON(w, http.StatusFound, listofbranches)

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))

	}
})
