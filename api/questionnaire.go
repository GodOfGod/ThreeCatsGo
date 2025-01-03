package api

import (
	"ThreeCatsGo/config"
	DB "ThreeCatsGo/database"
	"ThreeCatsGo/model"
	"ThreeCatsGo/tools"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QuestionnaireRouter(router *gin.Engine, db *sql.DB) {
	router.POST("/question/submit_questionnaire", ApiSubmitQuestionnaire(db))
	router.GET("/question/get_questionnaire_by_id", ApiGetQestionnaireById(db))
	router.GET("/question/get_questionnaire_by_date", ApiGetQuestionnaireByDate(db))
	router.GET("/question/get_all_questionnaire", ApiGetAllQuestionnaire(db))
	router.POST("/question/update_questionnaire", ApiUpdateQuestionnaire(db))
	router.GET("/question/get_questionnaire_config", ApiGetAllQuestionnaireConfig(db))
	router.GET("/question/get_questionnaire_config_by_id", ApiGetQuestionnaireConfigById(db))
	router.POST("/question/create_questionnaire", ApiCreateQuestionnaire(db))
	router.DELETE("/question/delete_questionnaire_by_id", ApiDeleteQuestionnaireById(db))
	router.DELETE("/question/delete_questionnaire_config_by_id", ApiDeleteQuestionnaireConfigById(db))
}

func ApiSubmitQuestionnaire(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var questionnaire model.Questionnaire
		ctx.ShouldBind(&questionnaire)
		questionnaire.Id = tools.GenerateId().String()
		err := DB.InsertQuestionnaire(db, questionnaire)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func ApiGetQestionnaireById(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := ctx.GetQuery("id")
		questionnaire := DB.GetQuestionnaireById(db, id)
		ctx.JSON(http.StatusOK, gin.H{"data": questionnaire})
	}
}

func ApiGetQuestionnaireByDate(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		date, _ := ctx.GetQuery("date")
		questionnaire := DB.GetQuestionnaireByDate(db, date)
		ctx.JSON(http.StatusOK, gin.H{"data": questionnaire})
	}
}

func ApiGetAllQuestionnaire(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		questionnaires := DB.GetAllQuestionnaire(db)
		ctx.JSON(http.StatusOK, gin.H{"data": questionnaires})
	}
}

func ApiUpdateQuestionnaire(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var questionnaire model.Questionnaire
		ctx.ShouldBind(&questionnaire)
		err := DB.UpdateQuestionnaire(db, questionnaire)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func ApiGetAllQuestionnaireConfig(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		customConfigList := DB.GetCustomConfigFields(db)
		questionnaireItems, err := config.ReadDefualteConfigFromJson()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed"})
			return
		}
		var questionnaireConfig model.QuestionnaireConfigList
		questionnaireConfig.CustomConfigList = customConfigList
		if customConfigList == nil {
			questionnaireConfig.CustomConfigList = []model.CustomConfigFields{}
		}
		questionnaireConfig.QuestionnaireItemList = questionnaireItems

		ctx.JSON(http.StatusOK, gin.H{"data": questionnaireConfig})
	}
}

func ApiGetQuestionnaireConfigById(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := ctx.GetQuery("config_id")
		customQuestionnaireFields, err := DB.GetQuestionnaireConfigById(db, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "no data"})
			return
		}
		questinonaireList, err := config.ReadDefualteConfigFromJson()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed"})
			return
		}
		var questionnaireConfig model.QuestionnaireConfig
		questionnaireConfig.CustomConfigFields = customQuestionnaireFields.ConfigFields
		if questionnaireConfig.CustomConfigFields == "" {
			questionnaireConfig.CustomConfigFields = "[]"
		}
		questionnaireConfig.QuestionnaireItemList = questinonaireList

		ctx.JSON(http.StatusOK, gin.H{"data": questionnaireConfig})
	}
}

func ApiCreateQuestionnaire(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var questionnaireConfig model.CustomConfigFields
		ctx.ShouldBind(&questionnaireConfig)
		questionnaireConfig.Id = tools.GenerateId().String()
		err := DB.InsertQuestionnaireConfig(db, questionnaireConfig)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed"})
			return
		}
		customQuestionnaireFields, err := DB.GetQuestionnaireConfigById(db, questionnaireConfig.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": customQuestionnaireFields})
	}
}

func ApiDeleteQuestionnaireById(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := ctx.GetQuery("id")
		err := DB.DeleteQuestionnaireById(db, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func ApiDeleteQuestionnaireConfigById(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := ctx.GetQuery("config_id")
		err := DB.DeleteQuestionnaireConfigById(db, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}
