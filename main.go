package main

import (
	"GoWebserver/router"
	"GoWebserver/services"
	"GoWebserver/utils"
	"log"
	"net/http"
)

func main() {
	log.Println("main app")

	var dbconn = utils.GetConnection
	services.SetDB(dbconn)
	var appRouter = router.CreateRouter

	log.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
}
