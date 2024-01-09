package main

import (
	"log"
	"net/http"

	"github.com/grahamuk2018/GoWebserver/services"
	"github.com/grahamuk2018/GoWebserver/utils"

	"github.com/grahamuk2018/GoWebserver/router"
)

func main() {
	log.Println("main app")

	var dbconn = utils.GetConnection()
	services.SetDB(dbconn)
	var appRouter = router.CreateRouter()

	log.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
}
