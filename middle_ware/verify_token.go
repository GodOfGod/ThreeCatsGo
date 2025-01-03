package middleware

import (
	"ThreeCatsGo/tools"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.Request.Header
		path := ctx.Request.URL.Path

		// 跳过鉴权
		freePath := []string{"/login", "/assets", "/question", "/wen"}
		for _, p := range freePath {
			reg := regexp.MustCompile("(.)*" + p + "(.)*")
			if reg.FindString(path) != "" {
				return
			}
		}

		accessToken := header.Get("Access_token")
		if accessToken != "" {
			tokenInfo, err := tools.VerifyToken(accessToken)
			if err != nil {
				fmt.Println(tools.ColoredStr("VerifyToken failed").Red())
				panic(err)
			}
			ctx.Set("userId", tokenInfo.UserId)
			return
		}

		ctx.JSON(http.StatusForbidden, gin.H{"message": "authorization failed"})
		ctx.Abort()
	}
}
