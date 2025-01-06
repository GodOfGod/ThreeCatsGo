package database

import (
	model "ThreeCatsGo/model"
	"database/sql"
	"fmt"
)

const INSERT_QUESTIONNAIRE = "INSERT INTO questionnaire (id, date, config_id, questionnaire) VALUES (?, ?, ?, ?)"

func InsertQuestionnaire(db *sql.DB, questionnaire model.Questionnaire) error {
	_, err := db.Exec(INSERT_QUESTIONNAIRE, questionnaire.Id, questionnaire.Date, questionnaire.ConfigId, questionnaire.Questionnaire)
	if err != nil {
		panic(err)
	}
	return err
}

func GetQuestionnaireById(db *sql.DB, id string) model.Questionnaire {
	var questionnaire model.Questionnaire
	err := db.QueryRow("SELECT * FROM questionnaire WHERE id = ?", id).Scan(&questionnaire.Id, &questionnaire.Date, &questionnaire.ConfigId, &questionnaire.Questionnaire)
	if err != nil {
		panic(err)
	}
	return questionnaire
}

func GetQuestionnaireByDate(db *sql.DB, date string) model.Questionnaire {
	var questionnaire model.Questionnaire
	err := db.QueryRow("SELECT * FROM questionnaire WHERE date = ?", date).Scan(&questionnaire.Id, &questionnaire.Date, &questionnaire.ConfigId, &questionnaire.Questionnaire)
	if err != nil {
		panic(err)
	}
	return questionnaire
}

func GetAllQuestionnaire(db *sql.DB) []model.Questionnaire {
	rows, err := db.Query("SELECT * FROM questionnaire")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var questionnaires []model.Questionnaire
	for rows.Next() {
		var questionnaire model.Questionnaire
		err := rows.Scan(&questionnaire.Id, &questionnaire.Date, &questionnaire.ConfigId, &questionnaire.Questionnaire)
		if err != nil {
			panic(err)
		}
		questionnaires = append(questionnaires, questionnaire)
	}
	return questionnaires
}

func UpdateQuestionnaire(db *sql.DB, questionnaire model.Questionnaire) error {
	_, err := db.Exec("UPDATE questionnaire SET questionnaire = ? WHERE id = ?", questionnaire.Questionnaire, questionnaire.Id)
	if err != nil {
		panic(err)
	}
	return err
}

func GetCustomConfigFields(db *sql.DB) []model.CustomConfigFields {
	rows, err := db.Query("SELECT * FROM questionnaire_config")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var configFields []model.CustomConfigFields
	for rows.Next() {
		var config model.CustomConfigFields
		err := rows.Scan(&config.Id, &config.ConfigFields, &config.Title)
		if err != nil {
			panic(err)
		}
		configFields = append(configFields, config)
	}
	return configFields
}

func GetQuestionnaireConfigById(db *sql.DB, id string) (model.CustomConfigFields, error) {
	var config model.CustomConfigFields
	err := db.QueryRow("SELECT * FROM questionnaire_config WHERE id =?", id).Scan(&config.Id, &config.ConfigFields, &config.Title)
	if err != nil {
		panic(err)
	}
	return config, err
}

func InsertQuestionnaireConfig(db *sql.DB, config model.CustomConfigFields) error {
	_, err := db.Exec("INSERT INTO questionnaire_config (id, config_fields, title) VALUES (?,?,?)", config.Id, config.ConfigFields, config.Title)
	if err != nil {
		panic(err)
	}
	return err
}

func DeleteQuestionnaireById(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM questionnaire WHERE id =?", id)
	if err != nil {
		panic(err)
	}
	return err
}

func DeleteQuestionnaireConfigById(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM questionnaire_config WHERE id =?", id)
	if err != nil {
		panic(err)
	}
	return err
}

func InsertQuestionnaireItem(db *sql.DB, questionnaireItem model.QuestinonaireItem) error {
	fmt.Println(questionnaireItem)
	_, err := db.Exec("INSERT INTO question_item (id, field, title, input_type, options) VALUES (?,?,?,?,?)", questionnaireItem.Id, questionnaireItem.Field, questionnaireItem.Title, questionnaireItem.InputType, questionnaireItem.Options)
	if err != nil {
		panic(err)
	}
	return err
}

func GetAllQuestionnaireItems(db *sql.DB) []model.QuestinonaireItem {
	rows, err := db.Query("SELECT * FROM question_item")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var questionnaireItems []model.QuestinonaireItem
	for rows.Next() {
		var questionnaireItem model.QuestinonaireItem
		err := rows.Scan(&questionnaireItem.Id, &questionnaireItem.Field, &questionnaireItem.Title, &questionnaireItem.InputType, &questionnaireItem.Options)
		if err != nil {
			panic(err)
		}
		questionnaireItems = append(questionnaireItems, questionnaireItem)
	}
	return questionnaireItems
}
