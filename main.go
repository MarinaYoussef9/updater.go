package main

import (
	
	"github.com/jinzhu/gorm"
    
    "updater/model"
	"updater/router"
)


func main() {

    var db *gorm.DB
    model.ConnectDB(db)
    server.StartServer()
}
