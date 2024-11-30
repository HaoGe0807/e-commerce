package utils

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/request"
	"github.com/pkg/errors"
)

const (
	WebType           = 1
	AppType           = 2                 // bff（webservice-esl、webservice-store） 默认使用此type
	JwtSecret         = "sunmi_esl_666\n" // jwt鉴权用的key。最后面的\n换行符不可缺少，否则会报错
	JwtAuthHeaderItem = "authorization"
)

type MyClaims struct {
	UserId  int64  `json:"user_id"`
	Token   string `json:"token"`
	AppType int64  `json:"app_type"`

	jwt.StandardClaims
}

// BuildJwtToken 生成token
func BuildJwtToken(claim MyClaims) (string, error) {
	var expireTime time.Time

	nowTime := time.Now()
	if claim.AppType == AppType { // 7d
		expireTime = nowTime.Add(time.Hour * 24 * 7)
	} else if claim.AppType == WebType { // 24h
		expireTime = nowTime.Add(time.Hour * 24)
	}

	claims := MyClaims{
		UserId: claim.UserId,
		Token:  claim.Token,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
		},
	}

	// conf.JwtSecret 为加密密钥
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(JwtSecret))
	return token, err
}

func ParseJwtToken(req *http.Request) (*MyClaims, error) {
	token, err := request.ParseFromRequest(req, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtSecret), nil
		}, request.WithClaims(&MyClaims{}))

	if err != nil {
		return nil, errors.New("Unauthorized access to this resource")
	}
	if !token.Valid {
		return nil, errors.New("Token is not valid")
	}

	myClaim, ok := token.Claims.(*MyClaims)
	if !ok {
		return nil, errors.New("Token is not valid")
	}
	return myClaim, nil
}
