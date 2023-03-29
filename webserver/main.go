package main

import (
	"log"
	"net/http"
	"webserver/pdb"
	"webserver/router"
	"webserver/services"

	"github.com/ava-labs/avalanchego/snow/networking/router"
)

func main() {
	log.Println("in main App")
	var dbconn = pdb.GetConnection()
	services.SetDB(dbconn)

	var appRouter = router.CreateRouter()

	log.Println("Listening on Port 8080")
	log.Fatal(http.ListenAndServe("8080", appRouter))

}
