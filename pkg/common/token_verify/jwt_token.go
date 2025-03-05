package token_verify

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gomodule/redigo/redis"
	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/utils"
)

type Claims struct {
	UserId   string
	Platform string
	jwt.RegisteredClaims
}

func BuildClaims(userId, platform string, ttl int64) Claims {
	now := time.Now()
	return Claims{
		UserId:   userId,
		Platform: platform,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(ttl*24) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}
}

// CreateToken 为用户创建 token，返回 token 过期时间和错误
func CreateToken(userId string, platformId int32) (string, int64, error) {
	claims := BuildClaims(userId, constant.PlatformIdToName(platformId), config.Config.TokenPolicy.AccessExpire)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.TokenPolicy.Secret))
	if err != nil {
		return "", 0, err
	}

	m, err := db.DB.GetTokenMapByUidPid(userId, platformId)
	if err != nil && err != redis.ErrNil {
		return "", 0, err
	}

	var deleteTokenKey []string
	for k := range m {
		deleteTokenKey = append(deleteTokenKey, k)
	}
	if len(deleteTokenKey) > 0 {
		err = db.DB.DeleteTokenByUidPid(userId, platformId, deleteTokenKey)
		if err != nil {
			return "", 0, err
		}
	}

	err = db.DB.AddTokenFlag(userId, platformId, tokenString, constant.NormalToken)
	if err != nil {
		return "", 0, err
	}
	return tokenString, claims.ExpiresAt.Time.Unix(), nil
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.TokenPolicy.Secret), nil
	}
}

func GetClaimsFromToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, &constant.ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, &constant.ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, &constant.ErrTokenNotValidYet
			} else {
				return nil, &constant.ErrTokenUnknown
			}
		} else {
			return nil, &constant.ErrTokenNotValidYet
		}
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, &constant.ErrTokenNotValidYet
}

func ParseToken(tokenString string) (*Claims, error) {
	claims, err := GetClaimsFromToken(tokenString)
	if err != nil {
		log.Error("token validate err ", err.Error())
		return nil, err
	}

	m, err := db.DB.GetTokenMapByUidPid(claims.UserId, constant.PlatformNameToId(claims.Platform))
	if err != nil {
		log.Error("get token from redis err ", err.Error())
		return nil, &constant.ErrTokenInvalid
	}
	if m == nil {
		log.Error("get token from redis err, m is nil")
		return nil, &constant.ErrTokenInvalid
	}

	if v, ok := m[tokenString]; ok {
		switch v {
		case constant.NormalToken:
			return claims, nil
		case constant.InvalidToken:
			return nil, &constant.ErrTokenInvalid
		case constant.KickedToken:
			return nil, &constant.ErrTokenInvalid
		case constant.ExpiredToken:
			return nil, &constant.ErrTokenExpired
		default:
			return nil, &constant.ErrTokenUnknown
		}
	}
	return nil, &constant.ErrTokenUnknown
}

func VerifyToken(token, userId string) (bool, error) {
	claims, err := ParseToken(token)
	if err != nil {
		return false, err
	}
	if claims.UserId != userId {
		return false, &constant.ErrTokenUnknown
	}
	return true, nil
}

func GetUserIdFromToken(token string) (bool, string) {
	claims, err := ParseToken(token)
	if err != nil {
		return false, ""
	}
	return true, claims.UserId
}

// CheckAccess 检查是否有访问权限
func CheckAccess(opUserId, ownerUserId string) bool {
	if utils.IsContainString(opUserId, config.Config.Admin.UserIds) {
		return true
	}
	return opUserId == ownerUserId
}
