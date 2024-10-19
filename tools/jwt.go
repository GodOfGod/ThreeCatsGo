package tools

import (
	config "ThreeCatsGo/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtUserClaims struct {
	UserId string
	jwt.RegisteredClaims
}


func GenerateToken(userId string) string {
	claims := JwtUserClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)), // 过期时间24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	strToken, err := token.SignedString([]byte(config.SECRET_KEY))
	if err != nil {
		fmt.Println(ColoredStr("Generate jwt failed").Red())
		panic(err)
	}
	return strToken
}

func VerifyToken(tokenStr string) (*JwtUserClaims, error) {
	user := JwtUserClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &user, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET_KEY), nil
	})

	if claims, ok := token.Claims.(*JwtUserClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}