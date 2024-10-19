package database

import (
	"database/sql"
	"fmt"

	tools "ThreeCatsGo/tools"

	_ "github.com/go-sql-driver/mysql"
)


type PetInfo struct {
	Nick_name string `json:"nick_name"`
	Birthday string `json:"birthday"`
	Gender int `json:"gender"`
	Avatar_path string `json:"avatar_path"`
	Pet_id string `json:"pet_id"`
}


func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:admin2004@(localhost:3306)/THREE_CATS")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil && db != nil {
			db.Close()
		}
	}()
	
	// ping database
	if err := db.Ping(); err != nil {
		panic(tools.ColoredStr("ping failed").Red())
	}
	// create user table
	{
		// CREATE_EVENT_TABLE
		// CREATE_USER_TABLE
		// CREATE_PETS_TABLE
		// _, err = db.Exec(CREATE_PETS_TABLE)
		// if err != nil {
		// 	fmt.Println(tools.ColoredStr("create table failed").Red(), err)
		// 	panic(err)
		// }
	}
	
	return db;
}




func (pet PetInfo) InsertPet(db *sql.DB) {
	_, err := db.Exec(INSERT_PET, pet.Pet_id, pet.Nick_name, pet.Gender, pet.Avatar_path, pet.Birthday)
	if err != nil {
		fmt.Println(err)
		panic(tools.ColoredStr("Insert pet failed").Red())
	}
}
