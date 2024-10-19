package database

import (
	"database/sql"

	model "ThreeCatsGo/model"
)

func InsertFile(db *sql.DB, fileInfo model.FileInfo) {
	const INSERT_FILE = `
		INSERT INTO file (id, file_name, file_path, event_id, pet_id, user_id) VALUES(?, ?, ?, ?, ?, ?);
	`
	_, err := db.Exec(INSERT_FILE, fileInfo.Id, fileInfo.FileName, fileInfo.FilePath, fileInfo.EventId, fileInfo.PetId, fileInfo.UserId)
	if err != nil {
		panic(err)
	}
}

func CreateEvent(db *sql.DB, eventId string, petId string, userId string) string {
	const INSERT_EVENT = `
		INSERT INTO events (id, pet_id, user_id, is_draft) VALUES (?, ?, ?, 1);
	`
	_, err := db.Exec(INSERT_EVENT, eventId, petId, userId)
	if err != nil {
		recover()
		return ""
	}
	return eventId
}

func SaveEvent(db *sql.DB, eventInfo model.EventInfo) error {
	const UPDATE_EVENT = `
		UPDATE events
			SET title = ?, image_path = ?, file_ids = ?, record = ?, date = ?, is_draft = 0, finished = 0
			WHERE id = ?;
	`
	_, err := db.Exec(UPDATE_EVENT, eventInfo.Title, eventInfo.ImagePath, eventInfo.FileIds, eventInfo.Record, eventInfo.Date, eventInfo.Id)
	return err
}

func GetEventListByDate(db *sql.DB, userId string, date string) ([]model.EventInfo, error) {
	const GET_EVENT = `
		SELECT id, title, image_path, file_ids, user_id, pet_id, record, date, finished from events
			WHERE user_id = ? AND date = ? AND is_draft = 0;
	`
	result, err := db.Query(GET_EVENT, userId, date)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var eventList []model.EventInfo
	for result.Next() {
		var eventRow model.EventInfo
		result.Scan(&eventRow.Id, &eventRow.Title, &eventRow.ImagePath, &eventRow.FileIds, &eventRow.UserId, &eventRow.PetId, &eventRow.Record, &eventRow.Date, &eventRow.Finished)
		eventList = append(eventList, eventRow)
	}
	return eventList, err
}

func GetEventById(db *sql.DB, event_id string) model.EventInfo {
	const GET_EVENT = `
		SELECT * from events
			WHERE id = ?;
	`
	result := db.QueryRow(GET_EVENT, event_id)

	var event_item model.EventInfo
	err := result.Scan(&event_item.Id, &event_item.Title, &event_item.ImagePath, &event_item.FileIds, &event_item.Record, &event_item.UserId, &event_item.PetId, &event_item.IsDraft, &event_item.Date, &event_item.Finished)
	if err != nil {
		panic(err)
	}
	return event_item
}

func GetFileById(db *sql.DB, fileIds string) []model.FileInfo {
	const GET_FILE_BY_ID = `
		select * from file where find_in_set(id, ?);
	`
	rows, _ := db.Query(GET_FILE_BY_ID, fileIds)
	var fileList []model.FileInfo
	for rows.Next() {
		var file model.FileInfo
		rows.Scan(&file.Id, &file.FileName, &file.FilePath, &file.EventId, &file.EventId, &file.UserId)
		fileList = append(fileList, file)
	}
	return fileList
}

func OperateEvent(db *sql.DB, event_id string, operate string, event_status int) {
	const FINISH_EVENT = `
		update events
			set finished = ?
		where id = ?;
	`
	const DELETE_EVENT = `
		delete from events
			where id = ?
	`
	if operate == "finish" {
		_, err := db.Exec(FINISH_EVENT, event_status, event_id)
		if err != nil {
			panic(err)
		}
		return
	}
	if operate == "delete" {
		_, err := db.Exec(DELETE_EVENT, event_id)
		if err != nil {
			panic(err)
		}
		return
	}
}

func GetEventByRange(db *sql.DB, userId string, begin_time string, end_time string) []model.EventInfo {
	const GET_EVENT_BY_RANGE = `
		select * from events
			where user_id = ? 
			and date between ? and ?;
	`
	rows, _ := db.Query(GET_EVENT_BY_RANGE, userId, begin_time, end_time)
	var event_list []model.EventInfo
	for rows.Next() {
		var event_item model.EventInfo
		rows.Scan(&event_item.Id, &event_item.Title, &event_item.ImagePath, &event_item.FileIds, &event_item.Record, &event_item.UserId, &event_item.PetId, &event_item.IsDraft, &event_item.Date, &event_item.Finished)
		event_list = append(event_list, event_item)
	}
	return event_list
}
