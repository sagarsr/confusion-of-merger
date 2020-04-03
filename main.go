package main

import (
	"bankapp/api"
	"bankapp/app"
	"bankapp/utils"
	"context"

	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

//Connection global
var Connection *pgx.Conn

func main() {
	log.Info("Database url is " + os.Getenv("DATABASE_URL"))
	port := os.Getenv("PORT")
	log.Info("Running on port " + port)
	Connection, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Error("error connecting to database!")
	}
	utils.SetConnection(Connection)
	defer Connection.Close(context.Background())
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/get-token", app.GetTokenHandler).Methods("GET")
	router.Handle("/banklist", app.ModJWTHandler(api.ViewListBanks)).Methods("GET")
	router.Handle("/branchlist", app.ModJWTHandler(api.ViewBranchList)).Methods("GET")

	//FrontEnd API Validation
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET"})

	// Launch server
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
