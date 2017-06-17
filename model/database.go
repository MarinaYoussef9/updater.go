package model
import(
		"log"

		"github.com/jinzhu/gorm"
       _"github.com/lib/pq"
)

type Update_Request struct{
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	Os 	 		string 	   
	Os_ver 		string
	Os_arch 	string 
	Vlc_ver 	string
}

func ConnectDB(db *gorm.DB){
	// TODO : Move the psqlinfo to config
	psqlInfo := "host=localhost dbname=updater user=postgres password=postgres sslmode=disable"
	db , _ = gorm.Open("postgres" , psqlInfo)
	defer db.Close()
  	db.AutoMigrate(&Update_Request{})
  	db.LogMode(true)
  	log.Println("Done with db")
}