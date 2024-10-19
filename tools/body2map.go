package tools

import (
	"encoding/json"
	"net/http"
)

func ReadBodyToMap(resp *http.Response)map[string]any {
	content := map[string]interface{}{}
	err := json.NewDecoder(resp.Body).Decode(&content)
	if err != nil {
		panic(ColoredStr("ReadBodyToMap failed").Red())
	}
	return content
}