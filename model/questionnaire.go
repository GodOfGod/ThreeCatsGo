package model

import "ThreeCatsGo/config"

type Questionnaire struct {
	Id            string `json:"id"`
	Date          string `json:"date"`
	ConfigId      string `json:"config_id"`
	Questionnaire string `json:"questionnaire"`
}

// ['name', 'age'] 单独定制的字段列表
type CustomConfigFields struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	ConfigFields string `json:"config_fields"`
}

// 用户页面，用于渲染问卷列表
type QuestionnaireConfig struct {
	// 问卷中的问题项
	QuestionnaireItemList []config.QuestinonaireItem `json:"questionnaire_item_list"`
	// 当前文件id对应的问卷所需要展示的字段
	CustomConfigFields string `json:"custom_config_fields"`
}

// 管理页面，返回问卷列表和单独定制的字段列表
type QuestionnaireConfigList struct {
	// 问卷中的问题项
	QuestionnaireItemList []config.QuestinonaireItem `json:"questionnaire_item_list"`
	// 自定义问卷
	CustomConfigList []CustomConfigFields `json:"custom_config_list"`
}
