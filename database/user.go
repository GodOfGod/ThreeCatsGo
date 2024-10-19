package database

import (
	tools "ThreeCatsGo/tools"
	"database/sql"
	"errors"
	"fmt"
)

type UserInfo struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Create_time string `json:"create_time"`
	Avatar_path string `json:"avatar_path"`
	Pets_id string `json:"pets_id"`
	Event_id string  `json:"event_id"`
} 


func (user UserInfo) InsertUser(db *sql.DB) int64 {
	result, err := db.Exec(INSERT_USER, user.Id, user.Username, user.Create_time, user.Avatar_path, user.Pets_id, user.Event_id)

	if err != nil {
		fmt.Println(err)
		panic(tools.ColoredStr("Insert user failed").Red())
	}
	
	userId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		panic(tools.ColoredStr("Get last user id failed").Red())
	}

	return userId
}

func (user *UserInfo) SelectUserById(db *sql.DB, uid string) error{
	
	row := db.QueryRow(SELECT_USER_ID, uid)
	err := row.Scan(&user.Id, &user.Username, &user.Create_time, &user.Avatar_path, &user.Pets_id, &user.Event_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		fmt.Println(tools.ColoredStr("SelectUserById failed").Red())
	}
	return err

}
func (user UserInfo) UpdatePets (db *sql.DB, uid string, petId string, opt string ) {
	if opt == "Add" {
		res, err := db.Exec(UPDATE_USER_PETS, petId, petId, uid)
		if err != nil {
			// id, _ := res.LastInsertId()
			fmt.Println(tools.ColoredStr("update failed").Red(), res)
			panic(err)
		}
		fmt.Println("ADD", uid, petId)
	} else if opt == "Delete" {

	}
}

