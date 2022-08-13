package main

import (
	"github.com/MuriloRibg/API-Rest_Gin.git/Database"
	"github.com/MuriloRibg/API-Rest_Gin.git/Routes"
)

func main() {
	Database.ConectarBanco()
	Routes.HandleRequests()
}
