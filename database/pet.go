package database

import "database/sql"

func SelectPetsById(db *sql.DB, idStr string) []PetInfo {
	rows, err := db.Query(SELECT_PETS_BY_ID, idStr)
	if err != nil {
		panic(err)
	}

	var petList []PetInfo
	for rows.Next() {
		var pet PetInfo
		rows.Scan(&pet.Pet_id, &pet.Nick_name, &pet.Birthday, &pet.Gender, &pet.Avatar_path)
		petList = append(petList, pet)
	}

	return petList
}
