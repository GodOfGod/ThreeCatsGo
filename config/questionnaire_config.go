package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type QuestinonaireItem struct {
	Field      string   `json:"field"`
	Title      string   `json:"title"`
	Required   bool     `json:"required"`
	InputType  string   `json:"inputType"`
	NeedRemark bool     `json:"needRemark"`
	Options    []string `json:"options,omitempty"` // Use omitempty for optional fields
}

func ReadDefualteConfigFromJson() ([]QuestinonaireItem, error) {
	// Open the JSON file
	file, err := os.Open("./config/questionnaire_config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file content
	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []QuestinonaireItem{}, err
	}

	var configList []QuestinonaireItem
	err = json.Unmarshal(byteValue, &configList)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return []QuestinonaireItem{}, err
	}

	return configList, nil
}
