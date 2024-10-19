package model

type FileInfo struct {
	Id       string `json:"id"`
	FilePath string `json:"file_path"`
	FileName string `json:"file_name"`
	UserId   string `json:"user_id"`
	PetId    string `json:"pet_id"`
	EventId  string `json:"event_id"`
}

type EventInfo struct {
	Id        string `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	ImagePath string `json:"image_path" db:"image_path"`
	FileIds   string `json:"file_ids" db:"file_ids"`
	UserId    string `json:"user_id" db:"user_id"`
	PetId     string `json:"pet_id" db:"pet_id"`
	IsDraft   int    `json:"is_draft" db:"is_draft"`
	Record    string `json:"record" db:"record"`
	Date      string `json:"date" db:"date"`
	Finished  int    `json:"finished" db:"finished"`
}
