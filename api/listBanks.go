package api

import (
	"bankapp/models"
	"bankapp/utils"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//ViewListBanks is the api endpoint
var ViewListBanks = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var branch models.Branch
	w.Header().Set("Content-Type", "application/json")
	ifsc, queryErr := r.URL.Query()["ifsc"]
	log.Info(ifsc)
	if !queryErr {
		utils.RespondWithError(w, http.StatusBadRequest, "Bank IFSC to be given")
		return

	}
	//Allowed method is GET
	// Can be extended with different HTTP methods
	switch r.Method {
	case "GET":
		err := branch.FetchBankDetailsWithIFSC(utils.Connection, ifsc[0])
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Sorry ! Unable to find data")
			return
		}
		utils.RespondWithJSON(w, http.StatusFound, branch)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))

	}
})
