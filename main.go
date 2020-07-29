package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tesarwijaya/ichos/config/database"
	"github.com/tesarwijaya/ichos/config/router"
	"github.com/tesarwijaya/ichos/middleware"
)

func main() {
	collector := database.InitCollector()

	r := router.Init(collector)
	r.Use(middleware.Logging)

	fmt.Println("serving at localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
