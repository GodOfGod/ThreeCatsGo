package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"ThreeCatsGo/config"
	DB "ThreeCatsGo/database"
	"ThreeCatsGo/model"
	"ThreeCatsGo/tools"

	"github.com/gin-gonic/gin"
)

func HandleRouter(router *gin.Engine, db *sql.DB) {
	router.GET("/get_event_list_by_id/:id", getEventListByUserId)
	router.GET("/get_event_list", getEventList)
	router.GET("/login", ApiLogin(db))
	router.POST("/add_pet", ApiAddPet(db))
	router.GET("/get_pets", ApiGetPetsById(db))
	router.GET("/get_user_info", ApiGetUserInfo(db))
	router.GET("/get_pets_list", ApiGetPetsList(db))
	router.POST("/add_event_img", ApiAddEventImg(db))
	router.POST("/add_event_file", ApiAddEventFile(db))
	router.GET("/create_event", ApiCreateEvent(db))
	router.POST("/save_event", ApiSaveEvent(db))
	router.GET("/get_event_list_by_date", ApiGetEventListByDate(db))
	router.GET("/get_event_by_id", ApiGetEventById(db))
	router.GET("/get_file_by_id", ApiGetFileById(db))
	router.GET("/operate_event", ApiEventOperate(db))
	router.GET("/get_event_by_range", ApiGetEventByRange(db))
}

func getEventListByUserId(ctx *gin.Context) {
}

func getEventList(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "test"})

}

func getUserInfo(db *sql.DB, id string) DB.UserInfo {
	var user DB.UserInfo
	err := (*DB.UserInfo).SelectUserById(&user, db, id)
	if err != nil {
		panic(err)
	}
	return user
}

func ApiGetUserInfo(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		userId, exist := ctx.Get("userId")
		fmt.Println(tools.ColoredStr(userId.(string)).Red(), "--------")
		if exist {
			user := getUserInfo(db, userId.(string))
			var pets []DB.PetInfo
			if user.Pets_id != "" {
				pets = DB.SelectPetsById(db, user.Pets_id)
			}

			resData := map[string]interface{}{
				"userInfo": user,
				"pets":     pets,
			}

			ctx.JSON(http.StatusOK, resData)
		} else {
			ctx.JSON(http.StatusForbidden, gin.H{"data": nil, "message": "未授权"})
		}
	}
}

func ApiLogin(db *sql.DB) func(ctx *gin.Context) {

	return func(ctx *gin.Context) {
		code, codeExist := ctx.GetQuery("code")
		username, _ := ctx.GetQuery("username")
		avatar_path, _ := ctx.GetQuery("avatar_path")
		session := getSessionByCode(code)
		user := DB.UserInfo{}

		// check user info
		if codeExist {
			user = getUserInfo(db, session.openid)

			// new user
			if user.Id == "" {
				userInfo := DB.UserInfo{
					Id:          session.openid,
					Username:    username,
					Create_time: time.Now().Format(time.DateOnly),
					Avatar_path: avatar_path,
					Pets_id:     "",
					Event_id:    "",
				}
				userInfo.InsertUser(db)
				user = userInfo
			}
			token := tools.GenerateToken(user.Id)
			ctx.Header("access_token", token)
			ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "登陆成功", "data": user})

		} else {
			ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "登陆失败", "data": nil})
		}
	}
}

// get openid from tencent as the unique id for a user
func getSessionByCode(code string) struct {
	session_key string
	openid      string
} {
	params := url.Values{}
	params.Add("appid", "wx8032f7dea79e58bf")
	params.Add("secret", "19f5d6207af4f059f456ec890cbc8f1e")
	params.Add("js_code", code)
	params.Add("grant_type", "authorization_code")
	params.Add("connect_redirect", "1")

	reqUrl := "https://api.weixin.qq.com/sns/jscode2session?" + params.Encode()
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		fmt.Println(err)
		panic(tools.ColoredStr("getSessionByCode").Red())
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(tools.ColoredStr("getSessionByCode client failed").Red())
	}

	defer resp.Body.Close()

	body := tools.ReadBodyToMap(resp)
	return struct {
		session_key string
		openid      string
	}{body["session_key"].(string), body["openid"].(string)}
}

func ApiAddPet(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// get parameters
		nickName, _ := ctx.GetPostForm("nick_name")
		genderStr, _ := ctx.GetPostForm("gender")
		gender, _ := strconv.Atoi(genderStr)
		birthday, _ := ctx.GetPostForm("birthday")
		avatarImg, _ := ctx.FormFile("avatar_img")
		// save image
		avatarImgPath := config.AVATAR_IMG_FOLDER + "/" + avatarImg.Filename
		ctx.SaveUploadedFile(avatarImg, avatarImgPath)

		petId := tools.GenerateId().String()
		// construct petinfo
		petAvatarResource := "http://localhost:8080/" + config.AVATAR_IMG_SOURCE + "/" + avatarImg.Filename
		petInfo := DB.PetInfo{Pet_id: petId, Nick_name: nickName, Gender: gender, Birthday: birthday, Avatar_path: petAvatarResource}

		// save in mysql pets
		petInfo.InsertPet(db)
		// update mysql users
		user := DB.UserInfo{}
		userId, exist := ctx.Get("userId")
		if exist {
			user.UpdatePets(db, userId.(string), petId, "Add")
		}

		ctx.JSON(http.StatusOK, gin.H{"data": petInfo, "code": 0, "message": "添加成功"})

		// ctx.Cookie()

	}
}

func ApiGetPetsById(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		petsId := ctx.Query("pets_id")
		pets := DB.SelectPetsById(db, petsId)
		ctx.JSON(http.StatusOK, gin.H{"data": pets})
	}
}

func ApiGetPetsList(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		userId, exist := ctx.Get("userId")
		if exist {
			user := getUserInfo(db, userId.(string))
			if user.Pets_id != "" {
				pets := DB.SelectPetsById(db, user.Pets_id)
				ctx.JSON(http.StatusOK, gin.H{"message": "ok", "data": pets})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "ok", "data": nil})
			return
		}
		ctx.JSON(http.StatusForbidden, gin.H{"message": "unauthorization", "data": nil})
	}
}

func ApiAddEventImg(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		img, _ := ctx.FormFile("event_img")
		saveImgPath := config.EVENT_IMG_FOLDER + "/" + img.Filename
		ctx.SaveUploadedFile(img, saveImgPath)
		imgPath := "http://localhost:8080/" + config.EVENT_IMG_SOURCE + "/" + img.Filename
		ctx.JSON(http.StatusOK, gin.H{"message": "ok", "data": imgPath})
	}
}

func ApiAddEventFile(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")
		fileName, _ := ctx.GetPostForm("file_name")
		petId, _ := ctx.GetPostForm("pet_id")
		eventId, _ := ctx.GetPostForm("event_id")
		userId, exist := ctx.Get("userId")
		if !exist {
			ctx.JSON(http.StatusForbidden, gin.H{"message": "unauthorization"})
			return
		}
		saveFilePath := config.EVENT_FILE_FOLDER + "/" + file.Filename
		ctx.SaveUploadedFile(file, saveFilePath)
		fileInfo := model.FileInfo{
			Id:       tools.GenerateId().String(),
			FilePath: "http://localhost:8080/" + config.EVENT_FILE_SOURCE + "/" + file.Filename,
			FileName: fileName,
			PetId:    petId,
			EventId:  eventId,
			UserId:   userId.(string),
		}
		DB.InsertFile(db, fileInfo)

		ctx.JSON(http.StatusOK, gin.H{"message": "ok", "data": fileInfo.Id})
	}
}

func ApiCreateEvent(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		pet_id, _ := ctx.GetQuery("pet_id")
		userId, _ := ctx.Get("userId")
		eventId := tools.GenerateId().String()
		DB.CreateEvent(db, eventId, pet_id, userId.(string))
		resData := map[string]interface{}{
			"event_id": eventId,
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok", "data": resData})
	}
}

func ApiSaveEvent(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		userId, _ := ctx.Get("userId")
		var eventInfo model.EventInfo
		ctx.ShouldBind(&eventInfo)
		eventInfo.UserId = userId.(string)
		err := DB.SaveEvent(db, eventInfo)
		if err != nil {
			recover()
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err, "data": nil})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func ApiGetEventListByDate(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		date, _ := ctx.GetQuery("date")
		userId, _ := ctx.Get("userId")
		eventList, err := DB.GetEventListByDate(db, userId.(string), date)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
			return
		}

		if eventList == nil {
			eventList = []model.EventInfo{}
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok", "data": eventList})
	}
}

func ApiGetEventById(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		event_id, _ := ctx.GetQuery("event_id")
		event_item := DB.GetEventById(db, event_id)
		ctx.JSON(http.StatusOK, gin.H{"data": event_item})
	}
}

func ApiGetFileById(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		file_ids, _ := ctx.GetQuery("file_ids")
		fileList := DB.GetFileById(db, file_ids)
		if fileList == nil {
			fileList = []model.FileInfo{}
		}
		ctx.JSON(http.StatusOK, gin.H{"data": fileList})
	}
}

func ApiEventOperate(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		event_id := ctx.Query("event_id")
		operate := ctx.Query("operate")
		event_status, _ := strconv.Atoi(ctx.Query("event_status"))
		DB.OperateEvent(db, event_id, operate, event_status)
		ctx.JSON(http.StatusOK, gin.H{"data": nil, "message": "ok"})
	}
}

func ApiGetEventByRange(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		begin_time := ctx.Query("begin_time")
		end_time := ctx.Query("end_time")
		userId, _ := ctx.Get("userId")

		event_list := DB.GetEventByRange(db, userId.(string), begin_time, end_time)
		ctx.JSON(http.StatusOK, gin.H{"data": event_list})

	}
}
